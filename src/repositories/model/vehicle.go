package model

import (
	"time"

	"github.com/caiiomp/vehicle-platform-sales/src/core/domain/entity"
)

type Vehicle struct {
	ID        string     `json:"id,omitempty" bson:"_id,omitempty"`
	VehicleID string     `json:"vehicle_id,omitempty" bson:"vehicle_id,omitempty"`
	Brand     string     `json:"brand,omitempty" bson:"brand,omitempty"`
	Model     string     `json:"model,omitempty" bson:"model,omitempty"`
	Year      int        `json:"year,omitempty" bson:"year,omitempty"`
	Color     string     `json:"color,omitempty" bson:"color,omitempty"`
	Price     float64    `json:"price,omitempty" bson:"price,omitempty"`
	UserID    string     `json:"user_id,omitempty" bson:"user_id,omitempty"`
	SoldAt    *time.Time `json:"sold_at,omitempty" bson:"sold_at,omitempty"`
	CreatedAt time.Time  `json:"created_at" bson:"created_at,omitempty"`
	UpdatedAt time.Time  `json:"updated_at" bson:"updated_at,omitempty"`
}

func VehicleFromDomain(vehicle entity.Vehicle) Vehicle {
	return Vehicle{
		ID:        vehicle.ID,
		VehicleID: vehicle.VehicleID,
		Brand:     vehicle.Brand,
		Model:     vehicle.Model,
		Year:      vehicle.Year,
		Color:     vehicle.Color,
		Price:     vehicle.Price,
		SoldAt:    vehicle.SoldAt,
	}
}

func (ref Vehicle) ToDomain() *entity.Vehicle {
	return &entity.Vehicle{
		ID:        ref.ID,
		Brand:     ref.Brand,
		Model:     ref.Model,
		Year:      ref.Year,
		Color:     ref.Color,
		Price:     ref.Price,
		SoldAt:    ref.SoldAt,
		CreatedAt: ref.CreatedAt,
		UpdatedAt: ref.UpdatedAt,
	}
}
