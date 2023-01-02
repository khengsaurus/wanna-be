package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func GetConnection() *sql.DB {
	dbName := os.Getenv("PG_DB_NAME")
	dbUser := os.Getenv("PG_USERNAME")
	dbPass := os.Getenv("PG_PASSWORD")
	dbHost := os.Getenv("PG_HOST")

	connStr := fmt.Sprintf(
		"postgres://%s:%s@%s/%s?sslmode=disable",
		dbUser, dbPass, dbHost, dbName,
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Print(err)
		return nil
	}

	return db
}
