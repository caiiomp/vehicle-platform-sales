package responses

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/caiiomp/vehicle-platform-sales/src/core/domain/entity"
)

func TestVehicleFromDomain(t *testing.T) {
	entityID := uuid.NewString()

	now := time.Now()

	vehicle := entity.Vehicle{
		ID:        1,
		EntityID:  entityID,
		Brand:     "Some Brand",
		Model:     "Some Model",
		Year:      2025,
		Color:     "Gray",
		Price:     80000,
		CreatedAt: now,
		UpdatedAt: now,
	}

	expected := Vehicle{
		ID:        1,
		EntityID:  entityID,
		Brand:     "Some Brand",
		Model:     "Some Model",
		Year:      2025,
		Color:     "Gray",
		Price:     80000,
		CreatedAt: now,
		UpdatedAt: now,
	}

	actual := VehicleFromDomain(vehicle)

	assert.Equal(t, expected, actual)
}
