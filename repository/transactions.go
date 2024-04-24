package repository

import (
	"database/sql"

	"github.com/redianmf/transportation-ticketing-app/domain"
)

func GetLastTransactionByUserId(db *sql.DB, user domain.User) (result domain.Transaction, err error) {
	sql := `SELECT * FROM transactions WHERE user_id = $1 ORDER BY created_at DESC LIMIT 1`

	rows, err := db.Query(sql, user.ID)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var transaction = domain.Transaction{}

		err = rows.Scan(&transaction.ID, &transaction.UserId, &transaction.TransportationModeId, &transaction.TransactionReference, &transaction.Type, &transaction.Status, &transaction.Amount, &transaction.StartPoint, &transaction.EndPoint, &transaction.CreatedAt, &transaction.UpdatedAt)
		if err != nil {
			panic(err)
		}

		result = transaction
	}

	return
}

func InsertTransaction(db *sql.DB, transaction domain.Transaction) (err error) {
	sql := `
	INSERT INTO transactions (user_id, transportation_mode_id, transaction_reference, type, status, amount, start_point, end_point, created_at, updated_at)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
	`

	errs := db.QueryRow(sql, transaction.UserId, transaction.TransportationModeId, transaction.TransactionReference, transaction.Type, transaction.Status, transaction.Amount, transaction.StartPoint, transaction.EndPoint, transaction.CreatedAt, transaction.UpdatedAt)

	return errs.Err()
}

func UpdateTransaction(db *sql.DB, transaction domain.Transaction) (err error) {
	sql := `
	UPDATE transactions
	SET status = $1, amount = $2, end_point = $3, updated_at = $4
	WHERE id = $5
	`

	errs := db.QueryRow(sql, transaction.Status, transaction.Amount, transaction.EndPoint, transaction.UpdatedAt, transaction.ID)

	return errs.Err()
}
