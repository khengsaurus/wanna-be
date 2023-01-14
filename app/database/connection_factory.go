package database

type IConnectionFactory interface {
	GetConnectionString() string
	NewConnection(string) (*PooledConnection, error)
}
