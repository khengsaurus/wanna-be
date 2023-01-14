package database

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/lib/pq"
)

type IPgConnectionFactory interface {
	NewConnection(string) (*PgPooledConnection, error)
}

type PgConnectionFactory struct {
	IConnectionFactory
	IPgConnectionFactory
}

func (c PgConnectionFactory) GetConnectionString() string {
	dbName := os.Getenv("PG_DB_NAME")
	dbUser := os.Getenv("PG_USERNAME")
	dbPass := os.Getenv("PG_PASSWORD")
	dbHost := os.Getenv("PG_HOST")

	return fmt.Sprintf(
		"postgres://%s:%s@%s/%s?sslmode=disable",
		dbUser, dbPass, dbHost, dbName,
	)
}

func (c PgConnectionFactory) NewConnection(ip string) (*PgPooledConnection, error) {
	connStr := c.GetConnectionString()
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	var conn PgPooledConnection
	conn.Key = ip
	conn.LastUsedAt = time.Now()
	conn.connection = db
	return &conn, nil
}
