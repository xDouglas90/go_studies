package db

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq"
)

func DbConnect() *sql.DB {
	dbUser := os.Getenv("POSTGRES_USER")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")

	conn := "user=" + dbUser + " password=" + dbPassword + " dbname=" + dbName + "host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conn)
	if err != nil {
		panic(err.Error())
	}

	return db
}
