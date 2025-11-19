package responses

import (
	"time"

	"github.com/caiiomp/vehicle-platform-sales/src/core/domain/entity"
)

type Sale struct {
	ID             string    `json:"id,omitempty"`
	VehicleID      string    `json:"vehicle_id"`
	DocumentNumber string    `json:"document_number"`
	Price          float64   `json:"price"`
	SoldAt         time.Time `json:"sold_at"`
}

func SaleFromDomain(sale entity.Sale) Sale {
	return Sale{
		ID:             sale.ID,
		VehicleID:      sale.VehicleID,
		DocumentNumber: sale.DocumentNumber,
		Price:          sale.Price,
		SoldAt:         sale.SoldAt,
	}
}
