package useCaseProducts

import (
	"context"

	uuid "github.com/satori/go.uuid"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

type ProductsRepository interface {
	CreateProduct(ctx context.Context, product entity.Product) error
	GetProducts(ctx context.Context, qFilters entity.QueryFilters) ([]entity.Product, error)
	GetProductByUid(ctx context.Context, uid uuid.UUID) (entity.Product, bool, error)
	GetProductsByNameLike(ctx context.Context, name string, qFilters entity.QueryFilters) ([]entity.Product, error)
	GetProductsByNameLikeWithExtra(ctx context.Context, name string, qFilters entity.QueryFilters) ([]entity.ProductWithExtra, error)
	GetProductsLikeNamesWithLimitOnEach(ctx context.Context, names []string, qFilters entity.QueryFilters) ([]entity.Product, error)
	GetProductsLikeNamesWithLimitOnEachWithExtra(ctx context.Context, names []string, qFilters entity.QueryFilters) ([]entity.ProductWithExtra, error)
	UpdateProduct(ctx context.Context, product entity.Product) error
	DeleteProduct(ctx context.Context, productUid uuid.UUID) error
	ReviveProduct(ctx context.Context, productUid uuid.UUID) error

	UpdateProductPhotos(ctx context.Context, uid uuid.UUID, imgPaths []string) error

	CreateProductCount(ctx context.Context, productUid uuid.UUID, count int64) error
	UpdateProductCount(ctx context.Context, productUid uuid.UUID, count int64) error
	GetProductCount(ctx context.Context, productUid uuid.UUID) (int64, bool, error)
	GetProductsWithExtra(ctx context.Context, qFilters entity.QueryFilters) ([]entity.ProductWithExtra, error)
	GetProductsTotal(ctx context.Context) (int64, error)
}
type CategoriesRepository interface {
	CreateCategory(ctx context.Context, category entity.Category) error
	GetCategoryByUid(ctx context.Context, uid uuid.UUID) (entity.Category, bool, error)
	GetCategoriesByNameLike(ctx context.Context, name string, qFilters entity.QueryFilters) ([]entity.Category, error)
	GetAllCategories(ctx context.Context) ([]entity.Category, error)
	UpdateCategory(ctx context.Context, category entity.Category) error
	DeleteCategory(ctx context.Context, uid uuid.UUID) error
	GetCategoriesByUserOrders(ctx context.Context, userUid uuid.UUID) ([]uuid.UUID, error)
}
