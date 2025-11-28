package model

import (
	"testing"
	"time"

	"github.com/caiiomp/vehicle-platform-sales/src/core/domain/entity"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestVehicleFromDomain(t *testing.T) {
	id := 1
	entityID := uuid.NewString()
	brand := uuid.NewString()
	model := uuid.NewString()
	year := 2025
	color := "Black"
	price := float64(95000)

	vehicle := entity.Vehicle{
		ID:       id,
		EntityID: entityID,
		Brand:    brand,
		Model:    model,
		Year:     year,
		Color:    color,
		Price:    price,
	}

	expected := Vehicle{
		ID:       id,
		EntityID: entityID,
		Brand:    brand,
		Model:    model,
		Year:     year,
		Color:    color,
		Price:    price,
	}

	actual := VehicleFromDomain(vehicle)

	assert.Equal(t, expected, actual)
}

func TestVehicleToDomain(t *testing.T) {
	id := 1
	entityID := uuid.NewString()
	brand := uuid.NewString()
	model := uuid.NewString()
	year := 2025
	color := "Black"
	price := float64(95000)
	now := time.Now()

	vehicle := Vehicle{
		ID:        id,
		EntityID:  entityID,
		Brand:     brand,
		Model:     model,
		Year:      year,
		Color:     color,
		Price:     price,
		CreatedAt: now,
		UpdatedAt: now,
	}

	expected := &entity.Vehicle{
		ID:        id,
		EntityID:  entityID,
		Brand:     brand,
		Model:     model,
		Year:      year,
		Color:     color,
		Price:     price,
		CreatedAt: now,
		UpdatedAt: now,
	}

	actual := vehicle.ToDomain()

	assert.Equal(t, expected, actual)
}
