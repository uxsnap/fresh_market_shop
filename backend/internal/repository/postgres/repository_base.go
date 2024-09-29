package repositoryPostgres

import (
	"context"
	"log"

	"github.com/pkg/errors"
	client "github.com/uxsnap/fresh_market_shop/backend/internal/client/database"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

type BasePgRepository struct {
	dbc client.ClientDB
}

func New(client client.ClientDB) *BasePgRepository {
	return &BasePgRepository{
		dbc: client,
	}
}

func (r *BasePgRepository) DB() client.DB {
	return r.dbc.DB()
}

func (r *BasePgRepository) WithTx(ctx context.Context, f func(ctx context.Context) error) (err error) {
	log.Printf("with tx")
	if !r.DB().HasTxInCtx(ctx) {
		var (
			tx         entity.Transaction
			beginTxErr error
		)
		ctx, tx, beginTxErr = r.DB().BeginTxWithContext(ctx)
		if beginTxErr != nil {
			return errors.Wrap(beginTxErr, "failed to start transaction")
		}

		defer func() {
			err = HandleTxEnd(ctx, tx, err)
		}()
	}

	return f(ctx)
}

func (r *BasePgRepository) BeginTxWithContext(ctx context.Context) (context.Context, entity.Transaction, error) {
	return r.dbc.DB().BeginTxWithContext(ctx)
}

func HandleTxEnd(ctx context.Context, tx entity.Transaction, err error) error {
	if err == nil {
		if commitErr := tx.Commit(ctx); commitErr != nil {
			log.Printf("with tx: failed to commit tx")
			return errors.Wrap(commitErr, "failed to commit transaction")
		}
		log.Printf("with tx: successfully commit tx")
		return nil
	}

	if rollbackErr := tx.Rollback(ctx); rollbackErr != nil {
		log.Printf("with tx: failed to rollback tx")
		return errors.Wrap(rollbackErr, "failed to rollback transaction")
	}
	log.Printf("with tx: successfully rollback tx")
	return nil
}
