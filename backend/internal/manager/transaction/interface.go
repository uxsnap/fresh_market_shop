package transaction

import (
	"context"

	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

type Transactor interface {
	BeginTxWithContext(ctx context.Context) (context.Context, entity.Transaction, error)
	HasTxInCtx(ctx context.Context) bool
}
