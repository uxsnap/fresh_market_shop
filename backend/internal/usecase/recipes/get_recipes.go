package useCaseRecipes

import (
	"context"
	"log"
	"time"

	"github.com/pkg/errors"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

func (uc *UseCaseRecipes) GetRecipes(
	ctx context.Context,
	cookingTime int64,
	createdAfter time.Time,
	limit uint64,
	offset uint64,
) ([]entity.Recipe, error) {
	log.Printf("ucRecipes.GetRecipes")

	recipes, err := uc.recipesRepository.GetRecipes(ctx, cookingTime, createdAfter, limit, offset)
	if err != nil {
		log.Printf("failed to get recipes: %v", err)
		return nil, errors.WithStack(err)
	}

	return recipes, nil
}
