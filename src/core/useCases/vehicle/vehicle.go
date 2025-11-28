package vehicle

import (
	"context"
	"errors"

	interfaces "github.com/caiiomp/vehicle-platform-sales/src/core/_interfaces"
	"github.com/caiiomp/vehicle-platform-sales/src/core/domain/entity"
)

type vehicleService struct {
	vehicleRepository              interfaces.VehicleRepository
	saleRepository                 interfaces.SaleRepository
	vehiclePlatformPaymentsAdapter interfaces.VehiclePlatformPaymentsAdapter
}

func NewVehicleService(
	vehicleRepository interfaces.VehicleRepository,
	saleRepository interfaces.SaleRepository,
	vehiclePlatformPaymentsAdapter interfaces.VehiclePlatformPaymentsAdapter,
) interfaces.VehicleService {
	return &vehicleService{
		vehicleRepository:              vehicleRepository,
		saleRepository:                 saleRepository,
		vehiclePlatformPaymentsAdapter: vehiclePlatformPaymentsAdapter,
	}
}

func (ref *vehicleService) Create(ctx context.Context, vehicle entity.Vehicle) (*entity.Vehicle, error) {
	return ref.vehicleRepository.Create(ctx, vehicle)
}

func (ref *vehicleService) GetByID(ctx context.Context, id string) (*entity.Vehicle, error) {
	return ref.vehicleRepository.GetByID(ctx, id)
}

func (ref *vehicleService) Search(ctx context.Context, isSold *bool) ([]entity.Vehicle, error) {
	return ref.vehicleRepository.Search(ctx, isSold)
}

func (ref *vehicleService) Update(ctx context.Context, id string, vehicle entity.Vehicle) (*entity.Vehicle, error) {
	current, err := ref.vehicleRepository.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if current == nil {
		return nil, nil
	}

	return ref.vehicleRepository.Update(ctx, id, vehicle)
}

func (ref *vehicleService) Buy(ctx context.Context, entityID, buyerDocumentNumber string) (*entity.Vehicle, error) {
	vehicle, err := ref.vehicleRepository.GetByID(ctx, entityID)
	if err != nil {
		return nil, err
	}

	if vehicle == nil {
		return nil, nil
	}

	existingSale, err := ref.saleRepository.GetByEntityID(ctx, entityID)
	if err != nil {
		return nil, err
	}

	if existingSale != nil && existingSale.Status == "APPROVED" {
		return nil, errors.New("vehicle already sold")
	}

	paymentID, err := ref.vehiclePlatformPaymentsAdapter.GeneratePayment(ctx, vehicle.Price, "APPROVED")
	if err != nil {
		return nil, err
	}

	sale := entity.Sale{
		EntityID:            entityID,
		PaymentID:           paymentID,
		BuyerDocumentNumber: buyerDocumentNumber,
		Price:               vehicle.Price,
	}

	_, err = ref.saleRepository.Create(ctx, sale)
	if err != nil {
		return nil, err
	}

	return vehicle, nil
}
