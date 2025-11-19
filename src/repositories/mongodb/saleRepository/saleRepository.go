package saleRepository

import (
	"context"

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
