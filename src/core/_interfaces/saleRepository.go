package interfaces

import (
	"context"
	"time"

	"github.com/caiiomp/vehicle-platform-sales/src/core/domain/entity"
)

type SaleRepository interface {
	Create(ctx context.Context, sale entity.Sale) (*entity.Sale, error)
	GetByEntityID(ctx context.Context, entityID string) (*entity.Sale, error)
	Search(ctx context.Context) ([]entity.Sale, error)
	UpdateStatusByPaymentID(ctx context.Context, paymentID, status string, soldDate time.Time) (*entity.Sale, error)
}
