package saleRepository

import (
	"context"
	"time"

	interfaces "github.com/caiiomp/vehicle-platform-sales/src/core/_interfaces"
	"github.com/caiiomp/vehicle-platform-sales/src/core/domain/entity"
	"github.com/caiiomp/vehicle-platform-sales/src/repositories/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type saleRepository struct {
	collection *mongo.Collection
}

func NewSaleRepository(collection *mongo.Collection) interfaces.SaleRepository {
	return &saleRepository{
		collection: collection,
	}
}

func (ref *saleRepository) Create(ctx context.Context, sale entity.Sale) (*entity.Sale, error) {
	record := model.SaleFromDomain(sale)

	result, err := ref.collection.InsertOne(ctx, record)
	if err != nil {
		return nil, err
	}

	objectID := result.InsertedID.(primitive.ObjectID)

	singleResult := ref.collection.FindOne(ctx, bson.M{"_id": objectID})

	var createdSale model.Sale
	if err = singleResult.Decode(&createdSale); err != nil {
		return nil, err
	}

	return createdSale.ToDomain(), nil
}

func (ref *saleRepository) Search(ctx context.Context) ([]entity.Sale, error) {
	cursor, err := ref.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	sales := make([]entity.Sale, 0)

	for cursor.Next(ctx) {
		var record model.Sale
		if err = cursor.Decode(&record); err != nil {
			return nil, err
		}

		sales = append(sales, *record.ToDomain())
	}

	return sales, nil
}

func (ref *saleRepository) GetByVehicleID(ctx context.Context, vehicleID string) (*entity.Sale, error) {
	result := ref.collection.FindOne(ctx, bson.M{"vehicle_id": vehicleID})
	if err := result.Err(); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	var record model.Sale
	if err := result.Decode(&record); err != nil {
		return nil, err
	}

	return record.ToDomain(), nil
}

func (ref *saleRepository) UpdateStatusByPaymentID(ctx context.Context, paymentID string, status string, soldDate *time.Time) (*entity.Sale, error) {
	update := bson.M{
		"$set": bson.M{
			"status":     status,
			"sold_at":    soldDate,
			"updated_at": time.Now(),
		},
	}

	_, err := ref.collection.UpdateOne(ctx, bson.M{"payment_id": paymentID}, update)
	if err != nil {
		return nil, err
	}

	result := ref.collection.FindOne(ctx, bson.M{"payment_id": paymentID})
	if err = result.Err(); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	var recordToReturn model.Sale
	if err = result.Decode(&recordToReturn); err != nil {
		return nil, err
	}

	return recordToReturn.ToDomain(), nil
}
