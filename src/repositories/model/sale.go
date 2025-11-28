package model

import (
	"time"

	"github.com/caiiomp/vehicle-platform-sales/src/core/domain/entity"
	valueobjects "github.com/caiiomp/vehicle-platform-sales/src/core/domain/valueObjects"
)

type Sale struct {
	ID                  int        `db:"id"`
	EntityID            string     `db:"entity_id"`
	PaymentID           string     `db:"payment_id"`
	BuyerDocumentNumber string     `db:"buyer_document_number"`
	Price               float64    `db:"price"`
	Status              string     `db:"status"`
	SoldAt              *time.Time `db:"sold_at"`
	CreatedAt           time.Time  `db:"created_at"`
	UpdatedAt           time.Time  `db:"created_at"`
}

func SaleFromDomain(sale entity.Sale) Sale {
	return Sale{
		EntityID:            sale.EntityID,
		PaymentID:           sale.PaymentID,
		BuyerDocumentNumber: sale.BuyerDocumentNumber,
		Price:               sale.Price,
		Status:              sale.Status.String(),
		SoldAt:              sale.SoldAt,
	}
}

func (ref *Sale) ToDomain() *entity.Sale {
	return &entity.Sale{
		ID:                  ref.ID,
		EntityID:            ref.EntityID,
		PaymentID:           ref.PaymentID,
		BuyerDocumentNumber: ref.BuyerDocumentNumber,
		Price:               ref.Price,
		Status:              valueobjects.SaleStatusType(ref.Status),
		SoldAt:              ref.SoldAt,
		CreatedAt:           ref.CreatedAt,
		UpdatedAt:           ref.UpdatedAt,
	}
}
