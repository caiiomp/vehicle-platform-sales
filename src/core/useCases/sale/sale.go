package sale

import (
	"context"
	"time"

	interfaces "github.com/caiiomp/vehicle-platform-sales/src/core/_interfaces"
	"github.com/caiiomp/vehicle-platform-sales/src/core/domain/entity"
)

type saleService struct {
	saleRepository interfaces.SaleRepository
	timeGenerator  func() *time.Time
}

func NewSaleService(saleRepository interfaces.SaleRepository, timeGenerator func() *time.Time) interfaces.SaleService {
	return &saleService{
		saleRepository: saleRepository,
		timeGenerator:  timeGenerator,
	}
}

func (ref *saleService) Create(ctx context.Context, sale entity.Sale) (*entity.Sale, error) {
	return ref.saleRepository.Create(ctx, sale)
}

func (ref *saleService) Search(ctx context.Context, status string) ([]entity.Sale, error) {
	return ref.saleRepository.Search(ctx, status)
}

func (ref *saleService) UpdateStatusByPaymentID(ctx context.Context, paymentID string, status string) (*entity.Sale, error) {
	var soldDate *time.Time
	if status == "APPROVED" {
		soldDate = ref.timeGenerator()
	}

	return ref.saleRepository.UpdateStatusByPaymentID(ctx, paymentID, status, soldDate)
}
