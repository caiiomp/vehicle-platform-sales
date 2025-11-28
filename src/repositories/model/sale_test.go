package model

import (
	"testing"
	"time"

	"github.com/caiiomp/vehicle-platform-sales/src/core/domain/entity"
	valueobjects "github.com/caiiomp/vehicle-platform-sales/src/core/domain/valueObjects"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestSaleFromDomain(t *testing.T) {
	entityID := uuid.NewString()
	paymentID := uuid.NewString()
	buyerDocumentNumber := uuid.NewString()
	price := float64(95000)
	status := "APPROVED"
	now := time.Now()

	sale := entity.Sale{
		EntityID:            entityID,
		PaymentID:           paymentID,
		BuyerDocumentNumber: buyerDocumentNumber,
		Price:               price,
		Status:              valueobjects.SaleStatusType(status),
		SoldAt:              &now,
	}

	expected := Sale{
		EntityID:            entityID,
		PaymentID:           paymentID,
		BuyerDocumentNumber: buyerDocumentNumber,
		Price:               price,
		Status:              status,
		SoldAt:              &now,
	}

	actual := SaleFromDomain(sale)

	assert.Equal(t, expected, actual)
}

func TestSaleToDomain(t *testing.T) {
	id := 1
	entityID := uuid.NewString()
	paymentID := uuid.NewString()
	buyerDocumentNumber := uuid.NewString()
	price := float64(95000)
	status := valueobjects.SaleStatusTypeApproved
	now := time.Now()
	yesterday := time.Now().Add(time.Hour * -24)

	sale := Sale{
		ID:                  id,
		EntityID:            entityID,
		PaymentID:           paymentID,
		BuyerDocumentNumber: buyerDocumentNumber,
		Price:               price,
		Status:              status.String(),
		SoldAt:              &now,
		CreatedAt:           yesterday,
		UpdatedAt:           yesterday,
	}

	expected := &entity.Sale{
		ID:                  id,
		EntityID:            entityID,
		PaymentID:           paymentID,
		BuyerDocumentNumber: buyerDocumentNumber,
		Price:               price,
		Status:              status,
		SoldAt:              &now,
		CreatedAt:           yesterday,
		UpdatedAt:           yesterday,
	}

	actual := sale.ToDomain()

	assert.Equal(t, expected, actual)
}
