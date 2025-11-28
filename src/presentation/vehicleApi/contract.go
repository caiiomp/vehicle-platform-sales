package vehicleApi

import (
	"github.com/caiiomp/vehicle-platform-sales/src/core/domain/entity"
)

type createVehicleRequest struct {
	VehicleID string  `json:"vehicle_id" binding:"required"`
	Brand     string  `json:"brand" binding:"required"`
	Model     string  `json:"model" binding:"required"`
	Year      int     `json:"year" binding:"required"`
	Color     string  `json:"color" binding:"required"`
	Price     float64 `json:"price" binding:"required"`
}

func (ref createVehicleRequest) ToDomain() *entity.Vehicle {
	return &entity.Vehicle{
		EntityID: ref.VehicleID,
		Brand:    ref.Brand,
		Model:    ref.Model,
		Year:     ref.Year,
		Color:    ref.Color,
		Price:    ref.Price,
	}
}

type entityUri struct {
	EntityID string `uri:"entity_id" binding:"required"`
}

type updateVehicleRequest struct {
	Brand string  `json:"brand"`
	Model string  `json:"model"`
	Year  int     `json:"year"`
	Color string  `json:"color"`
	Price float64 `json:"price"`
}

func (ref updateVehicleRequest) ToDomain() *entity.Vehicle {
	return &entity.Vehicle{
		Brand: ref.Brand,
		Model: ref.Model,
		Year:  ref.Year,
		Color: ref.Color,
		Price: ref.Price,
	}
}

type vehicleQuery struct {
	IsSold *bool `form:"is_sold"`
}

type buyVehicleRequest struct {
	BuyerDocumentNumber string `json:"buyer_document_number" binding:"required"`
}
