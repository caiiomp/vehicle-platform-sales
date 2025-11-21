package entity

import "time"

type Vehicle struct {
	ID        string
	VehicleID string
	Brand     string
	Model     string
	Year      int
	Color     string
	Price     float64
	CreatedAt time.Time
	UpdatedAt time.Time
}
