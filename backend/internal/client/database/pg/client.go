package pg

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
	client "github.com/uxsnap/fresh_market_shop/backend/internal/client/database"
)

type pgClient struct {
	dbc *pg
}

func NewClient(ctx context.Context, dsn string) (client.ClientDB, error) {
	pool, err := pgxpool.Connect(ctx, dsn)
	if err != nil {
		return nil, errors.Errorf("failed to connect to db: %v", err)
	}

	if err := pool.Ping(ctx); err != nil {
		return nil, errors.Errorf("failed to ping db: %v", err)
	}

	return &pgClient{
		dbc: &pg{pool: pool},
	}, nil
}

func (c *pgClient) DB() client.DB {
	return c.dbc
}

func (c *pgClient) Close(ctx context.Context) error {
	if c.dbc != nil {
		c.dbc.Close()
	}

	return nil
}
