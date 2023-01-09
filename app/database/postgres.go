package database

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/khengsaurus/wanna-be/consts"
	"github.com/khengsaurus/wanna-be/util"
	_ "github.com/lib/pq"
)

type Connection struct {
	key        string
	connection *sql.DB
	lastUsedAt time.Time
}

type ConnectionFactory interface {
	GetConnectionString() string
	NewConnection(ip string) (*Connection, error)
}

type PgConnectionFactory struct{}

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

func (c PgConnectionFactory) NewConnection(ip string) (*Connection, error) {
	connStr := c.GetConnectionString()
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	conn := &Connection{
		key:        ip,
		connection: db,
		lastUsedAt: time.Now(),
	}
	return conn, nil
}

func getClientConnKey(ip string) string {
	return fmt.Sprintf("client-%s", ip)
}

func GetConnFromReqCtx(r *http.Request) (*Connection, int) {
	connPool, ok := r.Context().Value(consts.PgConnPoolKey).(*ConnectionPool)
	if !ok {
		fmt.Println("Failed to access connection pool")
		return nil, http.StatusInternalServerError
	}

	ip := r.Header.Get("X-Forwarded-For")
	if ip == "" {
		return nil, http.StatusUnauthorized
	}

	conn, err := connPool.Connect(r.Context(), ip)
	if err != nil {
		fmt.Printf("%v\n", err)
		return nil, http.StatusInternalServerError
	}
	return conn, http.StatusOK
}

func (conn *Connection) Ping() error {
	return conn.connection.Ping()
}

func (conn *Connection) RunQuery(query string) ([]interface{}, error) {
	fmt.Printf("connection %s running %s\n", conn.key, query)
	db := conn.connection
	rows, err := db.Query(query)

	if err != nil {
		fmt.Printf("%v\n", err)
		return nil, err
	}

	data, err := util.ToJsonEncodable(rows)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (conn *Connection) Exec(query string, args ...any) (sql.Result, error) {
	return conn.connection.Exec(query, args...)
}
