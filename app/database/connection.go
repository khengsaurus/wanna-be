package database

import (
	"time"
)

type IConnection interface {
	Close() error
}

type IPooledConnection interface {
	Ping() error
}

type PooledConnection struct {
	IPooledConnection
	Key        string
	LastUsedAt time.Time
	Connection IConnection
}

type PooledConnectionRepr struct {
	ClientKey  string `json:"clientKey"`
	LastUsedAt int64  `json:"lastUsedAt"`
}
