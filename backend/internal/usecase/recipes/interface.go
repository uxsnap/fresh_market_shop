package useCaseRecipes

import (
	"context"

	uuid "github.com/satori/go.uuid"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

type ProductsRepository interface {
	GetProductsWithExtra(ctx context.Context, qFilters entity.QueryFilters) ([]entity.ProductWithExtra, error)
}

type RecipesRepository interface {
	CreateRecipe(ctx context.Context, recipe entity.Recipe) error
	GetRecipeByUid(ctx context.Context, uid uuid.UUID) (entity.Recipe, bool, error)
	GetRecipeSteps(ctx context.Context, uid uuid.UUID) ([]entity.RecipeStep, error)
	GetRecipesByNameLike(ctx context.Context, name string, qFilters entity.QueryFilters) ([]entity.Recipe, error)
	GetRecipes(ctx context.Context, qFilters entity.QueryFilters) ([]entity.Recipe, error)
	UpdateRecipe(ctx context.Context, recipe entity.Recipe) error
	DeleteRecipe(ctx context.Context, uid uuid.UUID) error

	DeleteRecipePhotos(ctx context.Context, uid uuid.UUID, photosUids ...uuid.UUID) error
}
