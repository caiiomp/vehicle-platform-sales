package responses

import (
	"testing"
	"time"

	"github.com/caiiomp/vehicle-platform-sales/src/core/domain/entity"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestSaleFromDomain(t *testing.T) {
	saleID := primitive.NewObjectID().Hex()
	vehicleID := primitive.NewObjectID().Hex()
	documentNumber := primitive.NewObjectID().Hex()

	now := time.Now()

	sale := entity.Sale{
		ID:             saleID,
		VehicleID:      vehicleID,
		DocumentNumber: documentNumber,
		Price:          80000,
		SoldAt:         now,
	}

	expected := Sale{
		ID:             saleID,
		VehicleID:      vehicleID,
		DocumentNumber: documentNumber,
		Price:          80000,
		SoldAt:         now,
	}

	actual := SaleFromDomain(sale)

	assert.Equal(t, expected, actual)
}
