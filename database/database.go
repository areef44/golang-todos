package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConnectDB() *sql.DB {
	dsn := "user=developer password=supersecretpassword host=localhost port=5432 dbname=todos sslmode=disable"
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}

	return db
}
