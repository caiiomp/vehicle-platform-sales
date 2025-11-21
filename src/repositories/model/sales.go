package model

import (
	"time"

	"github.com/caiiomp/vehicle-platform-sales/src/core/domain/entity"
)

type Sale struct {
	ID                  string     `json:"id,omitempty" bson:"_id,omitempty"`
	VehicleID           string     `json:"vehicle_id" bson:"vehicle_id"`
	PaymentID           string     `json:"payment_id" bson:"payment_id"`
	BuyerDocumentNumber string     `json:"document_number" bson:"buyer_document_number"`
	Price               float64    `json:"price" bson:"price"`
	Status              string     `json:"status" bson:"status"`
	SoldAt              *time.Time `json:"sold_at" bson:"sold_at"`
}

func SaleFromDomain(sale entity.Sale) Sale {
	return Sale{
		VehicleID:           sale.VehicleID,
		PaymentID:           sale.PaymentID,
		BuyerDocumentNumber: sale.BuyerDocumentNumber,
		Price:               sale.Price,
		Status:              sale.Status,
		SoldAt:              sale.SoldAt,
	}
}

func (ref *Sale) ToDomain() *entity.Sale {
	return &entity.Sale{
		ID:                  ref.ID,
		VehicleID:           ref.VehicleID,
		PaymentID:           ref.PaymentID,
		BuyerDocumentNumber: ref.BuyerDocumentNumber,
		Price:               ref.Price,
		Status:              ref.Status,
		SoldAt:              ref.SoldAt,
	}
}
