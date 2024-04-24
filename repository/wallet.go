package repository

import (
	"database/sql"

	"github.com/redianmf/transportation-ticketing-app/domain"
)

func GetWalletByUserId(db *sql.DB, wallet domain.Wallet) (result domain.Wallet, err error) {
	sql := `
	SELECT * FROM wallets WHERE user_id = $1 LIMIT 1
	`

	rows, err := db.Query(sql, wallet.UserId)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var wallet = domain.Wallet{}

		err = rows.Scan(&wallet.ID, &wallet.UserId, &wallet.Amount, &wallet.CreatedAt, &wallet.UpdatedAt)
		if err != nil {
			panic(err)
		}

		result = wallet
	}

	return
}

func InsertWallet(db *sql.DB, wallet domain.Wallet) (err error) {
	sql := `
	INSERT INTO wallets (user_id, amount, created_at, updated_at)
	VALUES ($1, $2, $3, $4)
	`
	errs := db.QueryRow(sql, wallet.UserId, wallet.Amount, wallet.CreatedAt, wallet.UpdatedAt)

	return errs.Err()
}

func UpdateWalletByUserId(db *sql.DB, wallet domain.Wallet) (err error) {
	sql := `
	UPDATE wallets
	SET amount = $1, updated_at = $2
	WHERE user_id = $3
	`

	errs := db.QueryRow(sql, wallet.Amount, wallet.UpdatedAt, wallet.UserId)

	return errs.Err()
}
