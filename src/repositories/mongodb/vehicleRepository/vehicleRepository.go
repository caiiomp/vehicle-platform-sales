package vehicleRepository

import (
	"context"
	"time"

	interfaces "github.com/caiiomp/vehicle-platform-sales/src/core/_interfaces"
	"github.com/caiiomp/vehicle-platform-sales/src/core/domain/entity"
	"github.com/caiiomp/vehicle-platform-sales/src/repositories/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type vehicleRepository struct {
	collection *mongo.Collection
}

func NewVehicleRepository(collection *mongo.Collection) interfaces.VehicleRepository {
	return &vehicleRepository{
		collection: collection,
	}
}

func (ref *vehicleRepository) Create(ctx context.Context, vehicle entity.Vehicle) (*entity.Vehicle, error) {
	record := model.VehicleFromDomain(vehicle)

	now := time.Now()
	record.CreatedAt = now
	record.UpdatedAt = now

	created, err := ref.collection.InsertOne(ctx, record)
	if err != nil {
		return nil, err
	}

	id := created.InsertedID.(primitive.ObjectID)

	result := ref.collection.FindOne(ctx, bson.M{"_id": id})
	if err = result.Err(); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	var recordToReturn model.Vehicle
	if err = result.Decode(&recordToReturn); err != nil {
		return nil, err
	}

	return recordToReturn.ToDomain(), nil
}

func (ref *vehicleRepository) GetByID(ctx context.Context, vehicleID string) (*entity.Vehicle, error) {
	result := ref.collection.FindOne(ctx, bson.M{"vehicle_id": vehicleID})
	if err := result.Err(); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	var record model.Vehicle
	if err := result.Decode(&record); err != nil {
		return nil, err
	}

	return record.ToDomain(), nil
}

func (ref *vehicleRepository) Search(ctx context.Context, isSold *bool) ([]entity.Vehicle, error) {
	filter := bson.M{}

	if isSold != nil {
		filter = bson.M{
			"sold_at": bson.M{"$eq": nil},
		}

		if *isSold {
			filter["sold_at"] = bson.M{"$ne": nil}
		}
	}

	sort := bson.D{{Key: "price", Value: 1}}

	findOptions := options.Find().SetSort(sort)

	cursor, err := ref.collection.Find(ctx, filter, findOptions)
	if err != nil {
		if err == mongo.ErrNilCursor {
			return nil, nil
		}
		return nil, err
	}

	records := make([]entity.Vehicle, 0)

	for cursor.Next(ctx) {
		var record model.Vehicle
		if err = cursor.Decode(&record); err != nil {
			return nil, err
		}

		records = append(records, *record.ToDomain())
	}

	return records, nil
}

func (ref *vehicleRepository) Update(ctx context.Context, vehicleID string, vehicle entity.Vehicle) (*entity.Vehicle, error) {
	record := model.VehicleFromDomain(vehicle)
	record.UpdatedAt = time.Now()

	update := bson.M{
		"$set": record,
	}

	_, err := ref.collection.UpdateOne(ctx, bson.M{"vehicle_id": vehicleID}, update)
	if err != nil {
		return nil, err
	}

	result := ref.collection.FindOne(ctx, bson.M{"vehicle_id": vehicleID})
	if err = result.Err(); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	var recordToReturn model.Vehicle
	if err = result.Decode(&recordToReturn); err != nil {
		return nil, err
	}

	return recordToReturn.ToDomain(), nil
}
