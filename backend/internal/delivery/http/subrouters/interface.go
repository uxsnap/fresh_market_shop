package subrouters

import (
	"context"
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

type ProductsService interface {
	CreateProduct(ctx context.Context, product entity.Product) (uuid.UUID, error)
	UpdateProduct(ctx context.Context, product entity.Product) error
	GetProductByUid(ctx context.Context, uid uuid.UUID) (entity.Product, bool, error)
	GetProducts(ctx context.Context, categoryUid uuid.UUID, ccalMin int64, ccalMax int64, createdBefore time.Time, createdAfter time.Time, limit uint64, offset uint64) ([]entity.Product, error)
	GetProductsWithCounts(ctx context.Context, categoryUid uuid.UUID, ccalMin int64, ccalMax int64, createdBefore time.Time, createdAfter time.Time, limit uint64, offset uint64) ([]entity.ProductWithStockQuantity, error)
	DeleteProduct(ctx context.Context, uid uuid.UUID) error

	UpdateProductCount(ctx context.Context, productUid uuid.UUID, stockQuantity int64) error
	IncrementProductCount(ctx context.Context, productUid uuid.UUID, incValue int64) error
	DecrementProductCount(ctx context.Context, productUid uuid.UUID, decValue int64) error
	GetProductCount(ctx context.Context, productUid uuid.UUID) (int64, bool, error)

	CreateCategory(ctx context.Context, category entity.Category) (uuid.UUID, error)
	GetCategoryByUid(ctx context.Context, uid uuid.UUID) (entity.Category, error)
	GetAllCategories(ctx context.Context) ([]entity.Category, error)
	UpdateCategory(ctx context.Context, category entity.Category) error
	DeleteCategory(ctx context.Context, uid uuid.UUID) error
}

type AuthService interface {
	Register(ctx context.Context, email string, password string) (uuid.UUID, error)
	UpdateUser(ctx context.Context, accessToken string, uid uuid.UUID, email string, password string) (accessJwt string, refreshJwt string, err error)
	GetUser(ctx context.Context, accessJwt string, uid uuid.UUID, email string) (entity.User, error)
	DeleteUser(ctx context.Context, accessJwt string, uid uuid.UUID) error
	Login(ctx context.Context, email string, password string) (accessJwt string, refreshJwt string, err error)
	Logout(ctx context.Context, accessJwt string, uid uuid.UUID) error
	Refresh(ctx context.Context, refreshToken string) (accessJwt string, refreshJwt string, err error)
	VerifyJwt(ctx context.Context, token string) error
	VerifyEmail(ctx context.Context, token string) error
	HealthCheck(ctx context.Context) entity.AuthServiceHealthCheck
}
