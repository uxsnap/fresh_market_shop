package DBclient

import (
	"context"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

type QueryExecer interface {
	Exec(ctx context.Context, sql string, args ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row

	ScanQueryRow(ctx context.Context, dest interface{}, sql string, args ...interface{}) error
	ScanAllQuery(ctx context.Context, dest interface{}, sql string, args ...interface{}) error
}

type Transactor interface {
	BeginTxWithContext(ctx context.Context) (context.Context, entity.Transaction, error)
}

type DB interface {
	QueryExecer
	Transactor
	HasTxInCtx(ctx context.Context) bool
	Ping(ctx context.Context) error
	Close()
}

type ClientDB interface {
	DB() DB
	Close(ctx context.Context) error
}
