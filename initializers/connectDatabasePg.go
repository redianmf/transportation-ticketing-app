package initializers

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
	"github.com/redianmf/transportation-ticketing-app/database"
)

var (
	DB *sql.DB
)

func ConnectDatabasePg() {
	var err error

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+"password=%s dbname=%s sslmode=disable", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

	DB, err = sql.Open("postgres", psqlInfo)
	err = DB.Ping()

	if err != nil {
		fmt.Println("DB Connection Failed")
	} else {
		fmt.Println("DB Connection Success")
	}

	database.DbMigrate(DB)
}
