package vehicle

import (
	"context"
	"errors"
	"time"

	interfaces "github.com/caiiomp/vehicle-platform-sales/src/core/_interfaces"
	"github.com/caiiomp/vehicle-platform-sales/src/core/domain/entity"
)

type vehicleService struct {
	vehicleRepository interfaces.VehicleRepository
	saleRepository    interfaces.SaleRepository
}

func NewVehicleService(vehicleRepository interfaces.VehicleRepository, saleRepository interfaces.SaleRepository) interfaces.VehicleService {
	return &vehicleService{
		vehicleRepository: vehicleRepository,
		saleRepository:    saleRepository,
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
	return ref.vehicleRepository.Update(ctx, id, vehicle)
}

func (ref *vehicleService) Buy(ctx context.Context, vehicleID, documentNumber string) (*entity.Vehicle, error) {
	vehicle, err := ref.vehicleRepository.GetByID(ctx, vehicleID)
	if err != nil {
		return nil, err
	}

	if vehicle == nil {
		return nil, errors.New("vehicle does not exist")
	}

	if vehicle.SoldAt != nil {
		return nil, errors.New("vehicle already sold")
	}

	soldTime := time.Now()

	sale := entity.Sale{
		VehicleID:      vehicleID,
		DocumentNumber: documentNumber,
		Price:          vehicle.Price,
		SoldAt:         soldTime,
	}

	_, err = ref.saleRepository.Create(ctx, sale)
	if err != nil {
		return nil, err
	}

	soldVehicle := entity.Vehicle{
		SoldAt: &soldTime,
	}

	return ref.vehicleRepository.Update(ctx, vehicleID, soldVehicle)
}
