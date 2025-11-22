package vehiclerepository

import (
	"context"
	"database/sql"

	interfaces "github.com/caiiomp/vehicle-platform-sales/src/core/_interfaces"
	"github.com/caiiomp/vehicle-platform-sales/src/core/domain/entity"
	"github.com/caiiomp/vehicle-platform-sales/src/repositories/model"
)

type vehicleRepository struct {
	db *sql.DB
}

func NewVehicleRepository(db *sql.DB) interfaces.VehicleRepository {
	return &vehicleRepository{
		db: db,
	}
}

func (ref *vehicleRepository) Create(ctx context.Context, vehicle entity.Vehicle) (*entity.Vehicle, error) {
	record := model.VehicleFromDomain(vehicle)

	row := ref.db.QueryRowContext(ctx, insertVehicle, record.EntityID, record.Brand, record.Model, record.Year, record.Color, record.Price)

	var created model.Vehicle
	err := row.Scan(&created.ID, &created.EntityID, &created.Brand, &created.Model, &created.Year, &created.Color, &created.Price, &created.CreatedAt, &created.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return created.ToDomain(), nil
}

func (ref *vehicleRepository) GetByID(ctx context.Context, id string) (*entity.Vehicle, error) {
	row := ref.db.QueryRowContext(ctx, getVehicleByEntityID, id)

	var vehicle model.Vehicle
	err := row.Scan(&vehicle.ID, &vehicle.EntityID, &vehicle.Brand, &vehicle.Model, &vehicle.Year, &vehicle.Color, &vehicle.Price, &vehicle.CreatedAt, &vehicle.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return vehicle.ToDomain(), nil
}

func (ref *vehicleRepository) Search(ctx context.Context, isSold *bool) ([]entity.Vehicle, error) {
	query := searchAllVehicles

	if isSold != nil {
		query = searchNotSoldVehicles
		if *isSold {
			query = searchSoldVehicles
		}
	}

	rows, err := ref.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	vehicles := make([]entity.Vehicle, 0)

	for rows.Next() {
		var record model.Vehicle
		err := rows.Scan(&record.ID, &record.EntityID, &record.Brand, &record.Model, &record.Year, &record.Color, &record.Price, &record.CreatedAt, &record.UpdatedAt)
		if err != nil {
			return nil, err
		}

		vehicles = append(vehicles, *record.ToDomain())
	}

	return vehicles, nil
}

func (ref *vehicleRepository) Update(ctx context.Context, id string, vehicle entity.Vehicle) (*entity.Vehicle, error) {
	row := ref.db.QueryRowContext(ctx, getVehicleByEntityID, id)

	var current model.Vehicle
	if err := row.Scan(&current.ID, &current.EntityID, &current.Brand, &current.Model, &current.Year, &current.Color, &current.Price, &current.CreatedAt, &current.UpdatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	var hasUpdate bool

	if vehicle.Brand != "" {
		current.Brand = vehicle.Brand
		hasUpdate = true
	}

	if vehicle.Model != "" {
		current.Model = vehicle.Model
		hasUpdate = true
	}

	if vehicle.Color != "" {
		current.Color = vehicle.Color
		hasUpdate = true
	}

	if vehicle.Year != 0 {
		current.Year = vehicle.Year
		hasUpdate = true
	}

	if vehicle.Price != 0 {
		current.Price = vehicle.Price
		hasUpdate = true
	}

	if !hasUpdate {
		return current.ToDomain(), nil
	}

	row = ref.db.QueryRowContext(ctx, updateVehicle, id, current.Brand, current.Model, current.Year, current.Color, current.Price)

	var updated model.Vehicle
	if err := row.Scan(&updated.ID, &updated.EntityID, &updated.Brand, &updated.Model, &updated.Year, &updated.Color, &updated.Price, &updated.CreatedAt, &updated.UpdatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return updated.ToDomain(), nil
}
