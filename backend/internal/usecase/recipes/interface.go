package useCaseRecipes

import (
	"context"
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

type RecipesRepository interface {
	CreateRecipe(ctx context.Context, recipe entity.Recipe) error
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
