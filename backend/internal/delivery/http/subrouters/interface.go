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
	GetProductsWithExtra(ctx context.Context, categoryUid uuid.UUID, ccalMin int64, ccalMax int64, createdBefore time.Time, createdAfter time.Time, limit uint64, offset uint64, withCounts bool, withPhotos bool, uuids []uuid.UUID) ([]entity.ProductWithExtra, error)
	GetProductsByNameLike(ctx context.Context, name string, limit uint64, offset uint64) ([]entity.Product, error)
	GetProductsByNameLikeWithExtra(ctx context.Context, name string, limit uint64, offset uint64, withCounts bool, withPhotos bool) ([]entity.ProductWithExtra, error)
	GetProductsLikeNamesWithLimitOnEach(ctx context.Context, names []string, limit uint64, offset uint64) ([]entity.Product, error)
	GetProductsLikeNamesWithLimitOnEachWithExtra(ctx context.Context, names []string, limit uint64, offset uint64, withCounts bool, withPhotos bool) ([]entity.ProductWithExtra, error)
	DeleteProduct(ctx context.Context, uid uuid.UUID) error

	UpdateProductCount(ctx context.Context, productUid uuid.UUID, stockQuantity int64) error
	IncrementProductCount(ctx context.Context, productUid uuid.UUID, incValue int64) error
	DecrementProductCount(ctx context.Context, productUid uuid.UUID, decValue int64) error
	GetProductCount(ctx context.Context, productUid uuid.UUID) (int64, bool, error)

	CreateCategory(ctx context.Context, category entity.Category) (uuid.UUID, error)
	GetCategoriesByNameLike(ctx context.Context, name string, limit uint64, offset uint64) ([]entity.Category, error)
	GetCategoryByUid(ctx context.Context, uid uuid.UUID) (entity.Category, error)
	GetAllCategories(ctx context.Context) ([]entity.Category, error)
	UpdateCategory(ctx context.Context, category entity.Category) error
	DeleteCategory(ctx context.Context, uid uuid.UUID) error
}

type AuthService interface {
	Register(ctx context.Context, email string, password string) (uuid.UUID, error)
	UpdateAuthUser(ctx context.Context, accessToken string, uid uuid.UUID, email string, password string) (accessJwt string, refreshJwt string, err error)
	GetAuthUser(ctx context.Context, accessJwt string, uid uuid.UUID, email string) (entity.AuthUser, error)
	DeleteAuthUser(ctx context.Context, accessJwt string, uid uuid.UUID) error
	Login(ctx context.Context, email string, password string) (accessJwt string, refreshJwt string, err error)
	Logout(ctx context.Context, accessJwt string, uid uuid.UUID) error
	Refresh(ctx context.Context, refreshToken string) (accessJwt string, refreshJwt string, err error)
	VerifyJwt(ctx context.Context, token string) error
	VerifyEmail(ctx context.Context, token string) error
	HealthCheck(ctx context.Context) entity.AuthServiceHealthCheck
}

type UsersService interface {
	CreateUser(ctx context.Context, user entity.User) (uuid.UUID, error)
	UpdateUser(ctx context.Context, user entity.User) error
	GetUser(ctx context.Context, uid uuid.UUID) (entity.User, bool, error)
	DeleteUser(ctx context.Context, uid uuid.UUID) error

	AddDeliveryAddress(ctx context.Context, address entity.DeliveryAddress) (uuid.UUID, error)
	UpdateDeliveryAddress(ctx context.Context, address entity.DeliveryAddress) error
	GetDeliveryAddress(ctx context.Context, uid uuid.UUID) (entity.DeliveryAddress, bool, error)
	GetUserDeliveryAddresses(ctx context.Context, userUid uuid.UUID) ([]entity.DeliveryAddress, error)
	DeleteDeliveryAddress(ctx context.Context, uid uuid.UUID) error
	DeleteUserDeliveryAddresses(ctx context.Context, userUid uuid.UUID) error
}

type RecipesService interface {
	CreateRecipe(ctx context.Context, recipe entity.Recipe) (uuid.UUID, error)
	GetRecipeByUid(ctx context.Context, uid uuid.UUID) (entity.Recipe, bool, error)
	GetRecipesByNameLike(ctx context.Context, name string, limit uint64, offset uint64) ([]entity.Recipe, error)
	GetRecipes(
		ctx context.Context,
		cookingTime int64,
		createdAfter time.Time,
		limit uint64,
		offset uint64,
	) ([]entity.Recipe, error)
	UpdateRecipe(ctx context.Context, recipe entity.Recipe) error
	DeleteRecipe(ctx context.Context, uid uuid.UUID) error
}

type OrdersService interface {
	CreateOrder(ctx context.Context, product entity.OrderProducts) (uuid.UUID, error)
}
