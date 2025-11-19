package sale

import (
	"context"
	"errors"
	"testing"
	"time"

	mocks "github.com/caiiomp/vehicle-platform-sales/src/core/_mocks"
	"github.com/caiiomp/vehicle-platform-sales/src/core/domain/entity"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestCreate(t *testing.T) {
	ctx := context.TODO()
	vehicleID := primitive.NewObjectID().Hex()
	documentNumber := primitive.NewObjectID().Hex()
	unexpectedError := errors.New("unexpected error")

	t.Run("should not create sale when failed to create", func(t *testing.T) {
		saleRepositoryMocked := mocks.NewSaleRepository(t)

		sale := entity.Sale{
			VehicleID:      vehicleID,
			DocumentNumber: documentNumber,
			Price:          50000,
			SoldAt:         time.Now(),
		}

		saleRepositoryMocked.On("Create", ctx, sale).
			Return(nil, unexpectedError)

		service := NewSaleService(saleRepositoryMocked)

		actual, err := service.Create(ctx, sale)

		assert.Nil(t, actual)
		assert.Equal(t, unexpectedError, err)
	})

	t.Run("should create sale successfully", func(t *testing.T) {
		saleRepositoryMocked := mocks.NewSaleRepository(t)

		sale := entity.Sale{
			VehicleID:      vehicleID,
			DocumentNumber: documentNumber,
			Price:          50000,
			SoldAt:         time.Now(),
		}

		saleRepositoryMocked.On("Create", ctx, sale).
			Return(&sale, nil)

		service := NewSaleService(saleRepositoryMocked)

		actual, err := service.Create(ctx, sale)

		assert.NotNil(t, actual)
		assert.Nil(t, err)
	})
}

func TestSearch(t *testing.T) {
	ctx := context.TODO()
	vehicleID := primitive.NewObjectID().Hex()
	documentNumber := primitive.NewObjectID().Hex()
	unexpectedError := errors.New("unexpected error")

	t.Run("should not search sales when failed to search", func(t *testing.T) {
		saleRepositoryMocked := mocks.NewSaleRepository(t)

		saleRepositoryMocked.On("Search", ctx).
			Return(nil, unexpectedError)

		service := NewSaleService(saleRepositoryMocked)

		actual, err := service.Search(ctx)

		assert.Nil(t, actual)
		assert.Equal(t, unexpectedError, err)
	})

	t.Run("should search sales successfully", func(t *testing.T) {
		saleRepositoryMocked := mocks.NewSaleRepository(t)

		sales := []entity.Sale{
			{
				VehicleID:      vehicleID,
				DocumentNumber: documentNumber,
				Price:          50000,
				SoldAt:         time.Now(),
			},
		}

		saleRepositoryMocked.On("Search", ctx).
			Return(sales, nil)

		service := NewSaleService(saleRepositoryMocked)

		actual, err := service.Search(ctx)

		assert.NotNil(t, actual)
		assert.Nil(t, err)
	})
}
