package useCaseRecipes

import (
	"context"
	"log"

	"github.com/pkg/errors"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

func (uc *UseCaseRecipes) GetRecipes(ctx context.Context, qFilters entity.QueryFilters) ([]entity.Recipe, error) {
	log.Printf("ucRecipes.GetRecipes")

	recipes, err := uc.recipesRepository.GetRecipes(ctx, qFilters)
	if err != nil {
		log.Printf("failed to get recipes: %v", err)
		return nil, errors.WithStack(err)
	}

	return recipes, nil
}
