package transaction

import (
	"context"
	"log"

	"github.com/pkg/errors"
	client "github.com/uxsnap/fresh_market_shop/backend/internal/client/database"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

type Manager struct {
	pgClient client.ClientDB
}

func NewTxManager(pgClient client.ClientDB) *Manager {
	return &Manager{
		pgClient: pgClient,
	}
}

type Tx struct {
	transactors []Transactor
}

func (m *Manager) NewPgTransaction() Tx {
	return m.NewTransaction(m.pgClient.DB())
}

func (m *Manager) NewTransaction(transactors ...Transactor) Tx {
	return Tx{transactors: transactors}
}

func (tx Tx) Execute(ctx context.Context, f func(ctx context.Context) error) (err error) {
	log.Printf("txManager: execute tx call")
	internalTransactions := make([]entity.Transaction, 0, len(tx.transactors))

	ctxTx := ctx
	for _, tr := range tx.transactors {
		var internalTx entity.Transaction

		ctxTx, internalTx, err = tr.BeginTxWithContext(ctx)
		if err != nil {
			return errors.WithStack(errors.Wrap(err, "failed to begin internal tx"))
		}

		internalTransactions = append(internalTransactions, internalTx)
	}

	defer func() {
		defer func() {
			tx.transactors = nil
		}()

		if r := recover(); r != nil {
			log.Printf("panic recovered")
			err = errors.Wrapf(err, "panic recovered: %v", r)
		}

		if err != nil {
			for idx, txn := range internalTransactions {
				if !tx.transactors[idx].HasTxInCtx(ctx) {

					if rollbackErr := txn.Rollback(ctxTx); rollbackErr != nil {
						err = errors.Wrapf(err, "rollback error: %v", rollbackErr)
					}

					log.Printf("rollback tx")
				}
			}
			return
		}

		for idx, txn := range internalTransactions {
			if !tx.transactors[idx].HasTxInCtx(ctx) {
				if commitErr := txn.Commit(ctxTx); commitErr != nil {
					err = errors.Wrapf(err, "commit error: %v", commitErr)
				}
				log.Printf("commit tx")
			}
		}
	}()

	if err := f(ctxTx); err != nil {
		return errors.WithStack(err)
	}

	return nil
}
