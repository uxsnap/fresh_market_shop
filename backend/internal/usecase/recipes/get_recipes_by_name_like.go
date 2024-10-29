package useCaseRecipes

import (
	"context"
	"log"

	"github.com/pkg/errors"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

func (uc *UseCaseRecipes) GetRecipesByNameLike(ctx context.Context, name string, qFilters entity.QueryFilters) ([]entity.Recipe, error) {
	log.Printf("ucRecipes.GetRecipesByNameLike: name %s", name)

	recipes, err := uc.recipesRepository.GetRecipesByNameLike(ctx, name, qFilters)
	if err != nil {
		log.Printf("failed to get recipes by name like %s: %v", name, err)
		return nil, errors.WithStack(err)
	}

	return recipes, nil
}
