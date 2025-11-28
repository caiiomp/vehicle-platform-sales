package responses

import (
	"time"

	"github.com/caiiomp/vehicle-platform-sales/src/core/domain/entity"
)

type Sale struct {
	ID                  int        `json:"id,omitempty"`
	VehicleID           string     `json:"vehicle_id"`
	PaymentID           string     `json:"payment_id"`
	BuyerDocumentNumber string     `json:"buyer_document_number"`
	Status              string     `json:"status"`
	Price               float64    `json:"price"`
	SoldAt              *time.Time `json:"sold_at,omitempty"`
}

func SaleFromDomain(sale entity.Sale) Sale {
	return Sale{
		ID:                  sale.ID,
		VehicleID:           sale.EntityID,
		PaymentID:           sale.PaymentID,
		BuyerDocumentNumber: sale.BuyerDocumentNumber,
		Status:              sale.Status.String(),
		Price:               sale.Price,
		SoldAt:              sale.SoldAt,
	}
}
