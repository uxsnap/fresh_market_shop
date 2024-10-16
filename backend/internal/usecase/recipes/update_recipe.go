package useCaseRecipes

import (
	"context"
	"log"
	"time"

	"github.com/pkg/errors"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

func (uc *UseCaseRecipes) UpdateRecipe(ctx context.Context, recipe entity.Recipe) error {
	log.Printf("ucRecipes.UpdateRecipe: uid %s", recipe.Uid)

	if err := validateRecipe(recipe); err != nil {
		log.Printf("failed to update recipe %s: %v", recipe.Uid, err)
		return errors.WithStack(err)
	}

	savedRecipe, isFound, err := uc.recipesRepository.GetRecipeByUid(ctx, recipe.Uid)
	if err != nil {
		log.Printf("failed to update recipe %s: %v", recipe.Uid, err)
		return errors.WithStack(err)
	}
	if !isFound {
		log.Printf("failed to update recipe %s: recipe not found", recipe.Uid)
		return errors.New("recipe not found")
	}

	recipe.CreatedAt = savedRecipe.CreatedAt
	recipe.UpdatedAt = time.Now().UTC()

	if err := uc.recipesRepository.UpdateRecipe(ctx, recipe); err != nil {
		log.Printf("failed to update recipe %s: %v", recipe.Uid, err)
		return errors.WithStack(err)
	}

	return nil
}
