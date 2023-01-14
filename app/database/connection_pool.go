package database

import (
	"context"
	"sync"
	"time"
)

type IConnectionPool interface {
	SetEvictionInterval()
	Connect(context.Context, string) (*PooledConnection, error)
	GetRepr() ConnectionPoolRepr
}

type ConnectionPool struct {
	IConnectionPool
	lock             sync.Mutex
	evictionInterval time.Duration
	connMaxIdle      time.Duration
	// connFactory      IConnectionFactory // override
	// connMap          map[string]*PooledConnection // override
}

type ConnectionPoolRepr struct {
	Total       int                    `json:"total"`
	Connections []PooledConnectionRepr `json:"connections"`
}
