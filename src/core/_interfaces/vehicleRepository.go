package interfaces

import (
	"context"

	"github.com/caiiomp/vehicle-platform-sales/src/core/domain/entity"
)

type VehicleRepository interface {
	Create(ctx context.Context, vehicle entity.Vehicle) (*entity.Vehicle, error)
	GetByID(ctx context.Context, id string) (*entity.Vehicle, error)
	Search(ctx context.Context, isSold *bool) ([]entity.Vehicle, error)
	Update(ctx context.Context, id string, vehicle entity.Vehicle) (*entity.Vehicle, error)
}
