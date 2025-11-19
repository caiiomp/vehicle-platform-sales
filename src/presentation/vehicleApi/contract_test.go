package vehicleApi

import (
	"testing"

	"github.com/caiiomp/vehicle-platform-sales/src/core/domain/entity"
	"github.com/stretchr/testify/assert"
)

func Test_createVehicleRequestToDomain(t *testing.T) {
	request := createVehicleRequest{
		Brand: "Some Brand",
		Model: "Some Model",
		Year:  2025,
		Color: "Gray",
		Price: 80000,
	}

	expected := &entity.Vehicle{
		Brand: "Some Brand",
		Model: "Some Model",
		Year:  2025,
		Color: "Gray",
		Price: 80000,
	}

	actual := request.ToDomain()

	assert.Equal(t, expected, actual)
}

func Test_updateVehicleRequestToDomain(t *testing.T) {
	request := updateVehicleRequest{
		Brand: "Some Brand",
		Model: "Some Model",
		Year:  2025,
		Color: "Gray",
		Price: 80000,
	}

	expected := &entity.Vehicle{
		Brand: "Some Brand",
		Model: "Some Model",
		Year:  2025,
		Color: "Gray",
		Price: 80000,
	}

	actual := request.ToDomain()

	assert.Equal(t, expected, actual)
}
