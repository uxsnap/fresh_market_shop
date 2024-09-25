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
}
