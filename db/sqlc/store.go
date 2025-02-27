package db

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

type Store interface {
	Querier
}
type SQLStore struct {
	connPool *pgxpool.Conn
	*Queries
}

func NewStore(connPool *pgxpool.Conn) *SQLStore {
	return &SQLStore{
		connPool: connPool,
		Queries:  New(connPool),
	}
}
