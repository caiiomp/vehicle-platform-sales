package responses

import (
	"testing"
	"time"

	"github.com/caiiomp/vehicle-platform-sales/src/core/domain/entity"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestSaleFromDomain(t *testing.T) {
	saleID := primitive.NewObjectID().Hex()
	vehicleID := primitive.NewObjectID().Hex()
	documentNumber := primitive.NewObjectID().Hex()
	paymentID := uuid.NewString()
	status := "APPROVED"

	now := time.Now()

	sale := entity.Sale{
		ID:                  saleID,
		VehicleID:           vehicleID,
		BuyerDocumentNumber: documentNumber,
		Price:               80000,
		SoldAt:              &now,
		PaymentID:           paymentID,
		Status:              status,
	}

	expected := Sale{
		ID:                  saleID,
		VehicleID:           vehicleID,
		BuyerDocumentNumber: documentNumber,
		Price:               80000,
		SoldAt:              &now,
		PaymentID:           paymentID,
		Status:              status,
	}

	actual := SaleFromDomain(sale)

	assert.Equal(t, expected, actual)
}
