package entity

import "time"

type Sale struct {
	ID                  int
	EntityID            string
	PaymentID           string
	BuyerDocumentNumber string
	Price               float64
	Status              string
	SoldAt              *time.Time
	CreatedAt           time.Time
	UpdatedAt           time.Time
}
