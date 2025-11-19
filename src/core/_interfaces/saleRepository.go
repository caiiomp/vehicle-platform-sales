package interfaces

import (
	"context"

	"github.com/caiiomp/vehicle-platform-sales/src/core/domain/entity"
)

type SaleRepository interface {
	Create(ctx context.Context, sale entity.Sale) (*entity.Sale, error)
	Search(ctx context.Context) ([]entity.Sale, error)
}
