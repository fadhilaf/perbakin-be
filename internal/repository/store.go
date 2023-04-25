package repository

import (
	postgresql "github.com/FadhilAF/perbakin-be/internal/repository/postgres/sqlc"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Store interface {
	postgresql.Querier
}

type storeImpl struct {
	postgresql.Querier
}

func NewStore(db *pgxpool.Pool) Store {
	return &storeImpl{
		Querier: postgresql.New(db),
	}
}
