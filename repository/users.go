package repository

import (
	"database/sql"

	"github.com/redianmf/transportation-ticketing-app/domain"
)

func InsertUser(db *sql.DB, user domain.User) (err error) {
	sql := `
	INSERT INTO users (username, email, password, birth_date, created_at, updated_at)
	VALUES ($1, $2, $3, $4, $5, $6)
	`
	errs := db.QueryRow(sql, user.Username, user.Email, user.Password, user.BirthDate, user.CreatedAt, user.UpdatedAt)

	return errs.Err()
}

func GetUserByEmail(db *sql.DB, user domain.User) (result domain.User, err error) {
	sql := `
	SELECT * FROM users WHERE users.email = $1 LIMIT 1
	`

	rows, err := db.Query(sql, user.Email)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var user = domain.User{}

		err = rows.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.BirthDate, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			panic(err)
		}

		result = user
	}

	return
}

func GetUserById(db *sql.DB, user domain.User) (result domain.User, err error) {
	sql := `
	SELECT * FROM users WHERE users.id = $1 LIMIT 1
	`

	rows, err := db.Query(sql, user.ID)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var user = domain.User{}

		err = rows.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.BirthDate, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			panic(err)
		}

		result = user
	}

	return
}
