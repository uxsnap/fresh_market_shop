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

	GetRecipesTotal(ctx context.Context) (int64, error)

	AddRecipeSteps(ctx context.Context, uid uuid.UUID, rSteps []entity.RecipeStep) error
	DeleteRecipeStep(ctx context.Context, uid uuid.UUID, step int) error
}
