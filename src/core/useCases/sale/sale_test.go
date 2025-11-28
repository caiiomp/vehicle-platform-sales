package sale

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"

	mocks "github.com/caiiomp/vehicle-platform-sales/src/core/_mocks"
	"github.com/caiiomp/vehicle-platform-sales/src/core/domain/entity"
	valueobjects "github.com/caiiomp/vehicle-platform-sales/src/core/domain/valueObjects"
)

func TestCreate(t *testing.T) {
	ctx := context.TODO()
	vehicleID := primitive.NewObjectID().Hex()
	documentNumber := primitive.NewObjectID().Hex()
	unexpectedError := errors.New("unexpected error")
	now := time.Now()

	t.Run("should not create sale when failed to create", func(t *testing.T) {
		saleRepositoryMocked := mocks.NewSaleRepository(t)

		sale := entity.Sale{
			EntityID:            vehicleID,
			BuyerDocumentNumber: documentNumber,
			Price:               50000,
			SoldAt:              &now,
		}

		saleRepositoryMocked.On("Create", ctx, sale).
			Return(nil, unexpectedError)

		service := NewSaleService(saleRepositoryMocked, time.Now)

		actual, err := service.Create(ctx, sale)

		assert.Nil(t, actual)
		assert.Equal(t, unexpectedError, err)
	})

	t.Run("should create sale successfully", func(t *testing.T) {
		saleRepositoryMocked := mocks.NewSaleRepository(t)

		sale := entity.Sale{
			EntityID:            vehicleID,
			BuyerDocumentNumber: documentNumber,
			Price:               50000,
			SoldAt:              &now,
		}

		saleRepositoryMocked.On("Create", ctx, sale).
			Return(&sale, nil)

		service := NewSaleService(saleRepositoryMocked, time.Now)

		actual, err := service.Create(ctx, sale)

		assert.NotNil(t, actual)
		assert.Nil(t, err)
	})
}

func TestSearch(t *testing.T) {
	ctx := context.TODO()
	entityID := uuid.NewString()
	documentNumber := primitive.NewObjectID().Hex()
	unexpectedError := errors.New("unexpected error")
	now := time.Now()

	t.Run("should not search sales when failed to search", func(t *testing.T) {
		saleRepositoryMocked := mocks.NewSaleRepository(t)

		saleRepositoryMocked.On("Search", ctx).
			Return(nil, unexpectedError)

		service := NewSaleService(saleRepositoryMocked, time.Now)

		actual, err := service.Search(ctx)

		assert.Nil(t, actual)
		assert.Equal(t, unexpectedError, err)
	})

	t.Run("should search sales successfully", func(t *testing.T) {
		saleRepositoryMocked := mocks.NewSaleRepository(t)

		sales := []entity.Sale{
			{
				EntityID:            entityID,
				BuyerDocumentNumber: documentNumber,
				Price:               50000,
				SoldAt:              &now,
			},
		}

		saleRepositoryMocked.On("Search", ctx).
			Return(sales, nil)

		service := NewSaleService(saleRepositoryMocked, time.Now)

		actual, err := service.Search(ctx)

		assert.NotNil(t, actual)
		assert.Nil(t, err)
	})
}

func TestUpdateStatusByPaymentID(t *testing.T) {
	ctx := context.TODO()
	vehicleID := uuid.NewString()
	paymentID := uuid.NewString()
	buyerDocumentNumber := uuid.NewString()
	status := valueobjects.SaleStatusTypeApproved
	soldAt := time.Now()

	sale := entity.Sale{
		ID:                  1,
		EntityID:            vehicleID,
		PaymentID:           paymentID,
		BuyerDocumentNumber: buyerDocumentNumber,
		Price:               50000,
		Status:              status,
		SoldAt:              &soldAt,
	}

	saleRepositoryMocked := mocks.NewSaleRepository(t)

	saleRepositoryMocked.On("UpdateStatusByPaymentID", ctx, paymentID, status.String(), mock.AnythingOfType("time.Time")).
		Return(&sale, nil)

	service := NewSaleService(saleRepositoryMocked, time.Now)

	expected := sale

	actual, err := service.UpdateStatusByPaymentID(ctx, paymentID, status.String())

	assert.Equal(t, &expected, actual)
	assert.Nil(t, err)
}
