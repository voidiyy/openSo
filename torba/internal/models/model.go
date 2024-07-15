package model

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

type Storage interface {
	UserStorage
}

type Postgres struct {
	DB *pgxpool.Pool
}
