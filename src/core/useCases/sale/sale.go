package sale

import (
	"context"

	interfaces "github.com/caiiomp/vehicle-platform-sales/src/core/_interfaces"
	"github.com/caiiomp/vehicle-platform-sales/src/core/domain/entity"
)

type saleService struct {
	saleRepository interfaces.SaleRepository
}

func NewSaleService(saleRepository interfaces.SaleRepository) interfaces.SaleService {
	return &saleService{
		saleRepository: saleRepository,
	}
}

func (ref *saleService) Create(ctx context.Context, sale entity.Sale) (*entity.Sale, error) {
	return ref.saleRepository.Create(ctx, sale)
}

func (ref *saleService) Search(ctx context.Context) ([]entity.Sale, error) {
	return ref.saleRepository.Search(ctx)
}
