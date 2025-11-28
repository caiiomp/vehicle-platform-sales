package interfaces

import (
	"context"
	"time"

	"github.com/caiiomp/vehicle-platform-sales/src/core/domain/entity"
)

type SaleRepository interface {
	Create(ctx context.Context, sale entity.Sale) (*entity.Sale, error)
	SearchByEntityID(ctx context.Context, entityID string) ([]entity.Sale, error)
	SearchByStatus(ctx context.Context, status string) ([]entity.Sale, error)
	UpdateStatusByPaymentID(ctx context.Context, paymentID, status string, soldDate *time.Time) (*entity.Sale, error)
}
