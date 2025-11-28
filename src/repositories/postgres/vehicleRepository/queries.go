package vehiclerepository

const (
	getVehicleByEntityID = "SELECT * FROM vehicles WHERE entity_id = $1;"

	insertVehicle = `
		INSERT INTO vehicles (entity_id, brand, model, year, color, price) 
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING *;
	`

	updateVehicle = `
		UPDATE vehicles
		SET
			brand = $2,
			model = $3,
			year = $4,
			color = $5,
			price = $6
		WHERE entity_id = $1
		RETURNING *;
	`

	searchAllVehicles = "SELECT * FROM vehicles ORDER BY price ASC;"

	searchSoldVehicles = `
		SELECT v.* FROM vehicles v
		JOIN sales s
		ON v.entity_id = s.entity_id
		WHERE s.status = 'APPROVED' and s.sold_at IS NOT NULL
		ORDER BY v.price ASC;
	`

	searchNotSoldVehicles = `
		SELECT v.* FROM vehicles v
		LEFT JOIN sales s
		ON v.entity_id = s.entity_id
		WHERE s.entity_id IS NULL OR (s.status != 'APPROVED' AND s.sold_at IS NULL)
		ORDER BY v.price ASC;
	`
)
