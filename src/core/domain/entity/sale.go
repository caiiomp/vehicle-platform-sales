package entity

import "time"

type Sale struct {
	ID                  string
	VehicleID           string
	PaymentID           string
	BuyerDocumentNumber string
	Price               float64
	Status              string
	SoldAt              *time.Time
}
