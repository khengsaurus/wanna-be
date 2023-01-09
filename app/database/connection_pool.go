package database

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type ConnectionPool struct {
	lock             sync.Mutex
	connFactory      ConnectionFactory
	connMap          map[string]*Connection
	evictionInterval time.Duration
	connMaxIdle      time.Duration
}

func NewConnectionPool(
	connBuilder ConnectionFactory,
	connMaxIdle time.Duration,
	evictionInterval time.Duration,
) *ConnectionPool {
	if evictionInterval == 0 {
		evictionInterval = time.Minute
	}

	pool := &ConnectionPool{
		connFactory:      connBuilder,
		connMap:          make(map[string]*Connection),
		evictionInterval: evictionInterval,
		connMaxIdle:      connMaxIdle,
	}

	go pool.evictConnectionPeriodically()

	return pool
}

// If there is an existing connection for the user, return it.
// Else create a new connection and return it.
func (pool *ConnectionPool) Connect(ctx context.Context, ip string) (*Connection, error) {
	pool.lock.Lock()
	defer pool.lock.Unlock()

	key := getClientConnKey(ip)
	fmt.Println(key)

	savedConn, ok := pool.connMap[key]
	if ok {
		savedConn.lastUsedAt = time.Now()
		return savedConn, nil
	}

	connection, err := pool.connFactory.NewConnection(ip)
	if err != nil {
		return nil, err
	}

	pool.connMap[key] = connection
	return connection, nil
}

func (pool *ConnectionPool) evictConnectionPeriodically() {
	ticker := time.NewTicker(pool.evictionInterval)
	defer ticker.Stop()

	for range ticker.C {
		pool.evictConnections()
	}
}

func (pool *ConnectionPool) evictConnections() {
	pool.lock.Lock()
	defer pool.lock.Unlock()

	for connInfo, connection := range pool.connMap {
		if time.Since(connection.lastUsedAt) > pool.connMaxIdle {
			fmt.Printf("Evicting connection %s", connInfo)

			if connection.connection != nil {
				_ = connection.connection.Close()
			}

			delete(pool.connMap, connInfo)
		}
	}
}
