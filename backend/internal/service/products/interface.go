package productsService

import (
	"context"

	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

type ProductsRepository interface {
	CreateProduct(ctx context.Context, product entity.Product) error
	//...
}
