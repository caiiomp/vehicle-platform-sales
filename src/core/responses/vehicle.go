package responses

import (
	"time"

	"github.com/caiiomp/vehicle-platform-sales/src/core/domain/entity"
)

type Vehicle struct {
	ID        string     `json:"id"`
	Brand     string     `json:"brand"`
	Model     string     `json:"model"`
	Year      int        `json:"year"`
	Color     string     `json:"color"`
	Price     float64    `json:"price"`
	SoldAt    *time.Time `json:"sold_at,omitempty"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

func VehicleFromDomain(vehicle entity.Vehicle) Vehicle {
	return Vehicle{
		ID:        vehicle.ID,
		Brand:     vehicle.Brand,
		Model:     vehicle.Model,
		Year:      vehicle.Year,
		Color:     vehicle.Color,
		Price:     vehicle.Price,
		SoldAt:    vehicle.SoldAt,
		CreatedAt: vehicle.CreatedAt,
		UpdatedAt: vehicle.UpdatedAt,
	}
}
