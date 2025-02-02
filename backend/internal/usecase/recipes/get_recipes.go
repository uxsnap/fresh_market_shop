package useCaseRecipes

import (
	"context"
	"log"

	"github.com/pkg/errors"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

func (uc *UseCaseRecipes) GetRecipes(ctx context.Context, qFilters entity.QueryFilters) (entity.RecipesWithTotal, error) {
	log.Printf("ucRecipes.GetRecipes")

	var res entity.RecipesWithTotal

	if err := uc.txManager.NewPgTransaction().Execute(ctx, func(ctx context.Context) error {
		recipes, err := uc.recipesRepository.GetRecipes(ctx, qFilters)
		if err != nil {
			log.Printf("failed to get recipes: %v", err)
			return errors.WithStack(err)
		}

		total, err := uc.recipesRepository.GetRecipesTotal(ctx)
		if err != nil {
			log.Printf("failed to get recipes total: %v", err)
			return errors.WithStack(err)
		}

		res.Recipes = recipes
		res.Total = total

		return nil
	}); err != nil {
		log.Printf("failed to get recipes: %v", err)
		return entity.RecipesWithTotal{}, errors.WithStack(err)
	}

	return res, nil
}
