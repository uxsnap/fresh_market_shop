package useCaseProducts

import (
	"context"

	uuid "github.com/satori/go.uuid"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

type ProductsRepository interface {
	CreateProduct(ctx context.Context, product entity.Product) error
	GetProducts(
		ctx context.Context,
		categoryUid uuid.UUID,
		ccalMin int64,
		ccalMax int64,
		limit uint64,
		offset uint64,
	) ([]entity.Product, error)
	GetProductsByCategory(ctx context.Context, categoryUid uuid.UUID, limit uint64, offset uint64) ([]entity.Product, error)
	GetProductByUid(ctx context.Context, uid uuid.UUID) (entity.Product, bool, error)
	UpdateProduct(ctx context.Context, product entity.Product) error
	DeleteProduct(ctx context.Context, productUid uuid.UUID) error

	CreateProductCount(ctx context.Context, productUid uuid.UUID, count int64) error
	UpdateProductCount(ctx context.Context, productUid uuid.UUID, count int64) error
	GetProductCount(ctx context.Context, productUid uuid.UUID) (int64, bool, error)
	GetProductsWithCounts(
		ctx context.Context,
		categoryUid uuid.UUID,
		ccalMin int64,
		ccalMax int64,
		limit uint64,
		offset uint64,
	) ([]entity.ProductWithStockQuantity, error)
}

type CategoriesRepository interface {
	GetAllCategories(ctx context.Context) ([]entity.Category, error)
}
