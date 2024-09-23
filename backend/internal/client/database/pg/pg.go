package pg

import (
	"context"
	"log"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

type TxKey struct{}

type pg struct {
	pool *pgxpool.Pool
}

func (p *pg) Exec(ctx context.Context, sql string, args ...interface{}) (pgconn.CommandTag, error) {
	if tx, ok := ctx.Value(TxKey{}).(pgx.Tx); ok {
		return tx.Exec(ctx, sql, args...)
	}

	return p.pool.Exec(ctx, sql, args...)
}

func (p *pg) Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error) {
	if tx, ok := ctx.Value(TxKey{}).(pgx.Tx); ok {
		return tx.Query(ctx, sql, args...)
	}

	return p.pool.Query(ctx, sql, args...)
}

func (p *pg) QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row {
	if tx, ok := ctx.Value(TxKey{}).(pgx.Tx); ok {
		return tx.QueryRow(ctx, sql, args...)
	}

	return p.pool.QueryRow(ctx, sql, args...)
}

func (p *pg) ScanQueryRow(ctx context.Context, dest interface{}, sql string, args ...interface{}) error {
	row, err := p.Query(ctx, sql, args...)
	if err != nil {
		return errors.WithStack(err)
	}

	return pgxscan.ScanOne(dest, row)
}

func (p *pg) ScanAllQuery(ctx context.Context, dest interface{}, sql string, args ...interface{}) error {
	rows, err := p.Query(ctx, sql, args...)
	if err != nil {
		return errors.WithStack(err)
	}

	return pgxscan.ScanAll(dest, rows)
}

func (p *pg) Ping(ctx context.Context) error {
	return p.pool.Ping(ctx)
}

func (p *pg) Close() {
	p.pool.Close()
}

func (p *pg) BeginTxWithContext(ctx context.Context) (context.Context, entity.Transaction, error) {
	if tx, ok := ctx.Value(TxKey{}).(pgx.Tx); ok {
		log.Printf("pg: tx already exist in ctx")
		return ctx, tx, nil
	}

	tx, err := p.pool.BeginTx(ctx, pgx.TxOptions{IsoLevel: pgx.ReadCommitted})
	if err != nil {
		log.Printf("failed to begin tx")
		return ctx, nil, errors.Wrap(err, "failed to begin tx")
	}

	log.Printf("begin new tx")
	return context.WithValue(ctx, TxKey{}, tx), tx, nil
}

func (p *pg) HasTxInCtx(ctx context.Context) bool {
	_, ok := ctx.Value(TxKey{}).(pgx.Tx)
	return ok
}
