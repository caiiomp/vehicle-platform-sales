package entity

import (
	"time"

	valueobjects "github.com/caiiomp/vehicle-platform-sales/src/core/domain/valueObjects"
)

type Sale struct {
	ID                  int
	EntityID            string
	PaymentID           string
	BuyerDocumentNumber string
	Price               float64
	Status              valueobjects.SaleStatusType
	SoldAt              *time.Time
	CreatedAt           time.Time
	UpdatedAt           time.Time
}
