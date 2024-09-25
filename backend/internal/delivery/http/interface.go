package deliveryHttp

import (
	"context"

	uuid "github.com/satori/go.uuid"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

type ProductsService interface {
	CreateProduct(ctx context.Context, product entity.Product) (uuid.UUID, error)
	UpdateProduct(ctx context.Context, product entity.Product) error
	GetProductByUid(ctx context.Context, uid uuid.UUID) (entity.Product, error)
	GetProductsWithPagination(ctx context.Context, limit, offset int) ([]entity.Product, error)
	DeleteProduct(ctx context.Context, uid uuid.UUID) error
}

type CategoriesService interface {
	GetAllCategories(ctx context.Context) ([]entity.Category, error)
}
