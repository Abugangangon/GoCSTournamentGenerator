package Database

import (
	"database/sql"
	_ "github.com/lib/pq"
)

const psqlInfo = "host=localhost port=5432 user=postgres password=postgres dbname=gctournament sslmode=disable"

func Connect() (*sql.DB, error) {
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	return db, err
}

func Close(db *sql.DB) {
	defer db.Close()
}