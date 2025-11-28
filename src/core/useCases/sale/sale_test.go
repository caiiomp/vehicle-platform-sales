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

		service := NewSaleService(saleRepositoryMocked, func() *time.Time { return &now })

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

		service := NewSaleService(saleRepositoryMocked, func() *time.Time { return &now })

		actual, err := service.Create(ctx, sale)

		assert.NotNil(t, actual)
		assert.Nil(t, err)
	})
}

func TestSearchByStatus(t *testing.T) {
	ctx := context.TODO()
	entityID := uuid.NewString()
	documentNumber := primitive.NewObjectID().Hex()
	unexpectedError := errors.New("unexpected error")
	now := time.Now()

	t.Run("should not search sales when failed to search", func(t *testing.T) {
		saleRepositoryMocked := mocks.NewSaleRepository(t)

		saleRepositoryMocked.On("SearchByStatus", ctx, "APPROVED").
			Return(nil, unexpectedError)

		service := NewSaleService(saleRepositoryMocked, func() *time.Time { return &now })

		actual, err := service.SearchByStatus(ctx, "APPROVED")

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

		saleRepositoryMocked.On("SearchByStatus", ctx, "APPROVED").
			Return(sales, nil)

		service := NewSaleService(saleRepositoryMocked, func() *time.Time { return &now })

		actual, err := service.SearchByStatus(ctx, "APPROVED")

		assert.NotNil(t, actual)
		assert.Nil(t, err)
	})
}

func TestUpdateStatusByPaymentID(t *testing.T) {
	ctx := context.TODO()
	vehicleID := uuid.NewString()
	paymentID := uuid.NewString()
	buyerDocumentNumber := uuid.NewString()
	now := time.Now()

	t.Run("should update payment status when its approved", func(t *testing.T) {
		sale := entity.Sale{
			ID:                  1,
			EntityID:            vehicleID,
			PaymentID:           paymentID,
			BuyerDocumentNumber: buyerDocumentNumber,
			Price:               50000,
			Status:              "APPROVED",
			SoldAt:              &now,
		}

		saleRepositoryMocked := mocks.NewSaleRepository(t)

		saleRepositoryMocked.On("UpdateStatusByPaymentID", ctx, paymentID, "APPROVED", mock.AnythingOfType("*time.Time")).
			Return(&sale, nil)

		service := NewSaleService(saleRepositoryMocked, func() *time.Time { return &now })

		expected := sale

		actual, err := service.UpdateStatusByPaymentID(ctx, paymentID, "APPROVED")

		assert.Equal(t, &expected, actual)
		assert.Nil(t, err)
	})

	t.Run("should update payment status when its not approved", func(t *testing.T) {
		var soldAt *time.Time
		sale := entity.Sale{
			ID:                  1,
			EntityID:            vehicleID,
			PaymentID:           paymentID,
			BuyerDocumentNumber: buyerDocumentNumber,
			Price:               50000,
			Status:              "CANCELED",
			SoldAt:              soldAt,
		}

		saleRepositoryMocked := mocks.NewSaleRepository(t)

		saleRepositoryMocked.On("UpdateStatusByPaymentID", ctx, paymentID, "CANCELED", soldAt).
			Return(&sale, nil)

		service := NewSaleService(saleRepositoryMocked, func() *time.Time { return &now })

		expected := sale

		actual, err := service.UpdateStatusByPaymentID(ctx, paymentID, "CANCELED")

		assert.Equal(t, &expected, actual)
		assert.Nil(t, err)
	})
}
