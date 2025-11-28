package salerepository

const (
	getSaleByEntityID = "SELECT * FROM sales WHERE entity_id = $1;"

	insertSale = `
		INSERT INTO sales (
			entity_id,
			payment_id,
			buyer_document_number,
			price,
			status,
			sold_at
		) 
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING *;
	`

	updateSaleStatusByPaymentID = `
		UPDATE sales SET 
			status = $2,
			sold_at = $3
		WHERE payment_id = $1 
		RETURNING *;
	`

	searchAllSales = "SELECT * FROM sales;"

	searchSalesByStatus = "SELECT * FROM sales WHERE status = $1;"
)
