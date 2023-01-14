package database

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/khengsaurus/wanna-be/util"
	_ "github.com/lib/pq"
)

type IPgPooledConnection interface {
	IPooledConnection
	RunQuery(string) ([]interface{}, error)
	Exec(string, ...any) (sql.Result, error)
}

type PgPooledConnection struct {
	PooledConnection
	IPgPooledConnection
	connection *sql.DB
}

func GetPgConnFromReq(r *http.Request) (*PgPooledConnection, int) {
	connPool, err := GetPgConnPoolFromReq(r)
	if err != nil {
		fmt.Printf("%v", err)
		return nil, http.StatusInternalServerError
	}

	ip := r.Header.Get("X-Forwarded-For")
	if ip == "" {
		return nil, http.StatusBadRequest
	}

	conn, err := connPool.Connect(r.Context(), ip)
	if err != nil {
		fmt.Printf("%v\n", err)
		return nil, http.StatusInternalServerError
	}
	return conn, http.StatusOK
}

func (conn *PgPooledConnection) Ping() error {
	return conn.connection.Ping()
}

func (conn *PgPooledConnection) RunQuery(query string) ([]interface{}, error) {
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

func (conn *PgPooledConnection) Exec(query string, args ...any) (sql.Result, error) {
	return conn.connection.Exec(query, args...)
}
