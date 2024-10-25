package repositoryProducts

import (
	"context"

	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
	errorWrapper "github.com/uxsnap/fresh_market_shop/backend/internal/error_wrapper"
)

func (r *ProductsRepository) UpdateCount(ctx context.Context, ops entity.OrderProducts) *errorWrapper.Error {
	return nil
}
