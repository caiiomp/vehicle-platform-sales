package salerepository

import (
	"context"
	"database/sql"
	"time"

	interfaces "github.com/caiiomp/vehicle-platform-sales/src/core/_interfaces"
	"github.com/caiiomp/vehicle-platform-sales/src/core/domain/entity"
	"github.com/caiiomp/vehicle-platform-sales/src/repositories/model"
)

type saleRepository struct {
	db *sql.DB
}

func NewSaleRepository(db *sql.DB) interfaces.SaleRepository {
	return &saleRepository{
		db: db,
	}
}

func (ref *saleRepository) Create(ctx context.Context, sale entity.Sale) (*entity.Sale, error) {
	record := model.SaleFromDomain(sale)

	row := ref.db.QueryRowContext(ctx, insertSale, record.EntityID, record.PaymentID, record.BuyerDocumentNumber, record.Price, record.Status, record.SoldAt)

	var created model.Sale
	err := row.Scan(&created.ID, &created.EntityID, &created.PaymentID, &created.BuyerDocumentNumber, &created.Price, &created.Status, &created.SoldAt, &created.CreatedAt, &created.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return created.ToDomain(), nil
}

func (ref *saleRepository) GetByEntityID(ctx context.Context, entityID string) (*entity.Sale, error) {
	row := ref.db.QueryRowContext(ctx, getSaleByEntityID, entityID)

	var sale model.Sale
	err := row.Scan(&sale.ID, &sale.EntityID, &sale.PaymentID, &sale.BuyerDocumentNumber, &sale.Price, &sale.Status, &sale.SoldAt, &sale.CreatedAt, &sale.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return sale.ToDomain(), nil
}

func (ref *saleRepository) Search(ctx context.Context, status string) ([]entity.Sale, error) {
	var (
		rows *sql.Rows
		err  error
	)

	switch {
	case status != "":
		rows, err = ref.db.QueryContext(ctx, searchSalesByStatus, status)
	default:
		rows, err = ref.db.QueryContext(ctx, searchAllSales)
	}

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	sales := make([]entity.Sale, 0)

	for rows.Next() {
		var record model.Sale
		err = rows.Scan(&record.ID, &record.EntityID, &record.PaymentID, &record.BuyerDocumentNumber, &record.Price, &record.Status, &record.SoldAt, &record.CreatedAt, &record.UpdatedAt)
		if err != nil {
			return nil, err
		}

		sales = append(sales, *record.ToDomain())
	}

	return sales, nil
}

func (ref *saleRepository) UpdateStatusByPaymentID(ctx context.Context, paymentID, status string, soldDate time.Time) (*entity.Sale, error) {
	row := ref.db.QueryRowContext(ctx, updateSaleStatusByPaymentID, paymentID, status, soldDate)

	var sale model.Sale
	err := row.Scan(&sale.ID, &sale.EntityID, &sale.PaymentID, &sale.BuyerDocumentNumber, &sale.Price, &sale.Status, &sale.SoldAt, &sale.CreatedAt, &sale.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return sale.ToDomain(), nil
}
