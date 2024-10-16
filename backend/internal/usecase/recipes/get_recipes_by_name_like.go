package useCaseRecipes

import (
	"context"
	"log"

	"github.com/pkg/errors"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

func (uc *UseCaseRecipes) GetRecipesByNameLike(ctx context.Context, name string, limit uint64, offset uint64) ([]entity.Recipe, error) {
	log.Printf("ucRecipes.GetRecipesByNameLike: name %s", name)

	recipes, err := uc.recipesRepository.GetRecipesByNameLike(ctx, name, limit, offset)
	if err != nil {
		log.Printf("failed to get recipes by name like %s: %v", name, err)
		return nil, errors.WithStack(err)
	}

	return recipes, nil
}
