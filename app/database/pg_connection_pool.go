package database

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/khengsaurus/wanna-be/consts"
	_ "github.com/lib/pq"
)

type PgConnectionPool struct {
	ConnectionPool
	connFactory PgConnectionFactory
	connMap     map[string]*PgPooledConnection
}

func NewPgConnectionPool(
	connBuilder PgConnectionFactory,
	connMaxIdle time.Duration,
	evictionInterval time.Duration,
) *PgConnectionPool {
	if evictionInterval == 0 {
		evictionInterval = time.Minute
	}

	var pgPool PgConnectionPool
	pgPool.connFactory = connBuilder
	pgPool.connMap = make(map[string]*PgPooledConnection)
	pgPool.evictionInterval = evictionInterval
	pgPool.connMaxIdle = connMaxIdle

	go pgPool.SetEvictionInterval()

	return &pgPool
}

func (pool *PgConnectionPool) SetEvictionInterval() {
	ticker := time.NewTicker(pool.evictionInterval)
	defer ticker.Stop()

	for range ticker.C {
		func() {
			pool.lock.Lock()
			defer pool.lock.Unlock()

			for connInfo, conn := range pool.connMap {
				if time.Since(conn.LastUsedAt) > pool.connMaxIdle {

					if conn.Connection != nil {
						_ = conn.Connection.Close()
					}

					delete(pool.connMap, connInfo)
				}
			}
		}()
	}
}

// If there is an existing connection for the user, return it.
// Else create a new connection and return it.
func (pool *PgConnectionPool) Connect(ctx context.Context, ip string) (*PgPooledConnection, error) {
	pool.lock.Lock()
	defer pool.lock.Unlock()

	key := fmt.Sprintf("pg-%s", ip)

	savedConn, ok := pool.connMap[key]
	if ok {
		savedConn.LastUsedAt = time.Now()
		return savedConn, nil
	}

	conn, err := pool.connFactory.NewConnection(ip)
	if err != nil {
		return nil, err
	}

	pool.connMap[key] = conn

	return conn, nil
}

func (pool *PgConnectionPool) GetRepr() ConnectionPoolRepr {
	total := len((*pool).connMap)
	connections := []PooledConnectionRepr{}

	for key, conn := range pool.connMap {
		connections = append(connections, PooledConnectionRepr{
			ClientKey:  key,
			LastUsedAt: conn.LastUsedAt.Unix(),
		})
	}

	return ConnectionPoolRepr{
		Total:       total,
		Connections: connections,
	}
}

func GetPgConnPoolFromReq(r *http.Request) (*PgConnectionPool, error) {
	connPool, ok := r.Context().Value(consts.PgConnPoolKey).(*PgConnectionPool)
	if !ok {
		return nil, fmt.Errorf("failed to access connection pool")
	}

	return connPool, nil
}
