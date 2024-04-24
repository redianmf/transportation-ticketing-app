package repository

import (
	"database/sql"

	"github.com/redianmf/transportation-ticketing-app/domain"
)

func GetAllTransportationMode(db *sql.DB) (results []domain.TransportationModes, err error) {
	sql := `SELECT * FROM transportation_modes`

	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var trMode = domain.TransportationModes{}

		err = rows.Scan(&trMode.ID, &trMode.Name, &trMode.BasePrice, &trMode.AdditionalPrice, &trMode.PriceCalculation, &trMode.CreatedAt, &trMode.UpdatedAt)
		if err != nil {
			panic(err)
		}

		results = append(results, trMode)
	}

	return
}

func GetTransportationModeById(db *sql.DB, transportationMode domain.TransportationModes) (result domain.TransportationModes, err error) {
	sql := `
	SELECT * FROM transportation_modes WHERE id = $1 LIMIT 1
	`

	rows, err := db.Query(sql, transportationMode.ID)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var trMode = domain.TransportationModes{}

		err = rows.Scan(&trMode.ID, &trMode.Name, &trMode.BasePrice, &trMode.AdditionalPrice, &trMode.PriceCalculation, &trMode.CreatedAt, &trMode.UpdatedAt)
		if err != nil {
			panic(err)
		}

		result = trMode
	}

	return
}

func InsertTransportationMode(db *sql.DB, transportationMode domain.TransportationModes) (err error) {
	sql := `
	INSERT INTO transportation_modes (name, base_price, additional_price, price_calculation, created_at, updated_at)
	VALUES ($1, $2, $3, $4, $5, $6)
	`

	errs := db.QueryRow(sql, transportationMode.Name, transportationMode.BasePrice, transportationMode.AdditionalPrice, transportationMode.PriceCalculation, transportationMode.CreatedAt, transportationMode.UpdatedAt)

	return errs.Err()
}

func UpdateTransportationMode(db *sql.DB, transportationMode domain.TransportationModes) (err error) {
	sql := `
	UPDATE transportation_modes
	SET name = $1, base_price = $2, additional_price = $3, price_calculation = $4, updated_at = $5
	WHERE id = $6
	`

	errs := db.QueryRow(sql, transportationMode.Name, transportationMode.BasePrice, transportationMode.AdditionalPrice, transportationMode.PriceCalculation, transportationMode.UpdatedAt, transportationMode.ID)

	return errs.Err()
}

func DeleteTransportationMode(db *sql.DB, transportationMode domain.TransportationModes) (err error) {
	sql := `DELETE FROM transportation_modes WHERE id = $1`

	errs := db.QueryRow(sql, transportationMode.ID)

	return errs.Err()
}
