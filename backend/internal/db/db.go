package db

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func New() (*pgxpool.Pool, error) {
	pool, err := pgxpool.New(context.Background(), os.Getenv("PG_DSN"))

	if err != nil {
		return nil, err

	}
	defer pool.Close()

	return pool, nil
}
