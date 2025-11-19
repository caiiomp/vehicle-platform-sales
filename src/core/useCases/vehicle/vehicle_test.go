package vehicle

import (
	"context"
	"errors"
	"testing"
	"time"

	mocks "github.com/caiiomp/vehicle-platform-sales/src/core/_mocks"
	"github.com/caiiomp/vehicle-platform-sales/src/core/domain/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestCreate(t *testing.T) {
	ctx := context.TODO()

	t.Run("should create vehicle successfully", func(t *testing.T) {
		vehicle := entity.Vehicle{
			Brand: "Some Brand",
			Model: "Some Model",
			Year:  2025,
			Color: "Gray",
			Price: 80000,
		}

		vehicleRepositoryMocked := mocks.NewVehicleRepository(t)

		vehicleRepositoryMocked.On("Create", ctx, vehicle).
			Return(&vehicle, nil)

		service := NewVehicleService(vehicleRepositoryMocked, nil)

		actual, err := service.Create(ctx, vehicle)

		assert.NotNil(t, actual)
		assert.Nil(t, err)
	})
}

func TestGetByID(t *testing.T) {
	ctx := context.TODO()
	vehicleID := primitive.NewObjectID().Hex()
	unexpectedError := errors.New("unexpected error")

	t.Run("should not get vehicle by id when failed to get", func(t *testing.T) {
		vehicleRepositoryMocked := mocks.NewVehicleRepository(t)

		vehicleRepositoryMocked.On("GetByID", ctx, vehicleID).
			Return(nil, unexpectedError)

		service := NewVehicleService(vehicleRepositoryMocked, nil)

		actual, err := service.GetByID(ctx, vehicleID)

		assert.Nil(t, actual)
		assert.Equal(t, unexpectedError, err)
	})

	t.Run("should get vehicle by id successfully", func(t *testing.T) {
		vehicleRepositoryMocked := mocks.NewVehicleRepository(t)

		now := time.Now()

		vehicle := &entity.Vehicle{
			ID:        vehicleID,
			Brand:     "Some Brand",
			Model:     "Some Model",
			Year:      2025,
			Color:     "Gray",
			Price:     80000,
			CreatedAt: now,
			UpdatedAt: now,
		}

		vehicleRepositoryMocked.On("GetByID", ctx, vehicleID).
			Return(vehicle, nil)

		service := NewVehicleService(vehicleRepositoryMocked, nil)

		actual, err := service.GetByID(ctx, vehicleID)

		assert.NotNil(t, actual)
		assert.Nil(t, err)
	})
}

func TestSearch(t *testing.T) {
	ctx := context.TODO()
	unexpectedError := errors.New("unexpected error")

	t.Run("should not search vehicles when failed to search", func(t *testing.T) {
		vehicleRepositoryMocked := mocks.NewVehicleRepository(t)

		isSold := true

		vehicleRepositoryMocked.On("Search", ctx, &isSold).
			Return(nil, unexpectedError)

		service := NewVehicleService(vehicleRepositoryMocked, nil)

		actual, err := service.Search(ctx, &isSold)

		assert.Nil(t, actual)
		assert.Equal(t, unexpectedError, err)
	})

	t.Run("should search vehicles successfully", func(t *testing.T) {
		vehicleRepositoryMocked := mocks.NewVehicleRepository(t)

		isSold := true

		vehicleRepositoryMocked.On("Search", ctx, &isSold).
			Return([]entity.Vehicle{}, nil)

		service := NewVehicleService(vehicleRepositoryMocked, nil)

		actual, err := service.Search(ctx, &isSold)

		assert.NotNil(t, actual)
		assert.Nil(t, err)
	})
}

func TestUpdate(t *testing.T) {
	ctx := context.TODO()
	vehicleID := primitive.NewObjectID().Hex()
	unexpectedError := errors.New("unexpected error")

	t.Run("should not update vehicle when failed to update", func(t *testing.T) {
		vehicleRepositoryMocked := mocks.NewVehicleRepository(t)

		vehicleRepositoryMocked.On("Update", ctx, vehicleID, entity.Vehicle{}).
			Return(nil, unexpectedError)

		service := NewVehicleService(vehicleRepositoryMocked, nil)

		actual, err := service.Update(ctx, vehicleID, entity.Vehicle{})

		assert.Nil(t, actual)
		assert.Equal(t, unexpectedError, err)
	})

	t.Run("should update vehicle successfully", func(t *testing.T) {
		vehicleRepositoryMocked := mocks.NewVehicleRepository(t)

		vehicleRepositoryMocked.On("Update", ctx, vehicleID, entity.Vehicle{}).
			Return(&entity.Vehicle{}, nil)

		service := NewVehicleService(vehicleRepositoryMocked, nil)

		actual, err := service.Update(ctx, vehicleID, entity.Vehicle{})

		assert.NotNil(t, actual)
		assert.Nil(t, err)
	})
}

func TestBuy(t *testing.T) {
	ctx := context.TODO()
	vehicleID := primitive.NewObjectID().Hex()
	userID := primitive.NewObjectID().Hex()
	unexpectedError := errors.New("unexpected error")

	t.Run("should not buy vehicle when failed to get vehicle by id", func(t *testing.T) {
		vehicleRepositoryMocked := mocks.NewVehicleRepository(t)
		saleRepositoryMocked := mocks.NewSaleRepository(t)

		vehicleRepositoryMocked.On("GetByID", ctx, vehicleID).
			Return(nil, unexpectedError)

		service := NewVehicleService(vehicleRepositoryMocked, nil)

		actual, err := service.Buy(ctx, vehicleID, userID)

		assert.Nil(t, actual)
		assert.Equal(t, unexpectedError, err)
		vehicleRepositoryMocked.AssertNumberOfCalls(t, "Update", 0)
		saleRepositoryMocked.AssertNumberOfCalls(t, "Create", 0)
	})

	t.Run("should not buy vehicle when vehicle does not exist", func(t *testing.T) {
		vehicleRepositoryMocked := mocks.NewVehicleRepository(t)
		saleRepositoryMocked := mocks.NewSaleRepository(t)

		vehicleRepositoryMocked.On("GetByID", ctx, vehicleID).
			Return(nil, nil)

		service := NewVehicleService(vehicleRepositoryMocked, nil)

		actual, err := service.Buy(ctx, vehicleID, userID)

		assert.Nil(t, actual)
		assert.ErrorContains(t, err, "vehicle does not exist")
		vehicleRepositoryMocked.AssertNumberOfCalls(t, "Update", 0)
		saleRepositoryMocked.AssertNumberOfCalls(t, "Create", 0)
	})

	t.Run("should not buy vehicle when vehicle already sold", func(t *testing.T) {
		vehicleRepositoryMocked := mocks.NewVehicleRepository(t)
		saleRepositoryMocked := mocks.NewSaleRepository(t)

		now := time.Now()

		vehicleAlreadySold := &entity.Vehicle{
			SoldAt: &now,
		}

		vehicleRepositoryMocked.On("GetByID", ctx, vehicleID).
			Return(vehicleAlreadySold, nil)

		service := NewVehicleService(vehicleRepositoryMocked, nil)

		actual, err := service.Buy(ctx, vehicleID, userID)

		assert.Nil(t, actual)
		assert.ErrorContains(t, err, "vehicle already sold")
		vehicleRepositoryMocked.AssertNumberOfCalls(t, "Update", 0)
		saleRepositoryMocked.AssertNumberOfCalls(t, "Create", 0)
	})

	t.Run("should not buy vehicle when failed to create sale", func(t *testing.T) {
		vehicleRepositoryMocked := mocks.NewVehicleRepository(t)
		saleRepositoryMocked := mocks.NewSaleRepository(t)

		vehicle := &entity.Vehicle{}

		vehicleRepositoryMocked.On("GetByID", ctx, vehicleID).
			Return(vehicle, nil)

		saleRepositoryMocked.On("Create", ctx, mock.AnythingOfType("entity.Sale")).
			Return(nil, unexpectedError)

		service := NewVehicleService(vehicleRepositoryMocked, saleRepositoryMocked)

		actual, err := service.Buy(ctx, vehicleID, userID)

		assert.Nil(t, actual)
		assert.Equal(t, unexpectedError, err)
		vehicleRepositoryMocked.AssertNumberOfCalls(t, "Update", 0)
	})

	t.Run("should buy vehicle successfully", func(t *testing.T) {
		vehicleRepositoryMocked := mocks.NewVehicleRepository(t)
		saleRepositoryMocked := mocks.NewSaleRepository(t)

		vehicle := &entity.Vehicle{}

		vehicleRepositoryMocked.On("GetByID", ctx, vehicleID).
			Return(vehicle, nil)
		vehicleRepositoryMocked.On("Update", ctx, vehicleID, mock.AnythingOfType("entity.Vehicle")).
			Return(vehicle, nil)

		saleRepositoryMocked.On("Create", ctx, mock.AnythingOfType("entity.Sale")).
			Return(nil, nil)

		service := NewVehicleService(vehicleRepositoryMocked, saleRepositoryMocked)

		actual, err := service.Buy(ctx, vehicleID, userID)

		assert.NotNil(t, actual)
		assert.Nil(t, err)
	})
}
