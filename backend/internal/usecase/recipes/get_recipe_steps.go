package useCaseRecipes

import (
	"context"
	"log"

	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

func (uc *UseCaseRecipes) GetRecipeSteps(ctx context.Context, uid uuid.UUID) ([]entity.RecipeStep, error) {
	log.Printf("ucRecipes.GetRecipeByUid: uid %s", uid)

	recipes, err := uc.recipesRepository.GetRecipeSteps(ctx, uid)

	if err != nil {
		log.Printf("failed to get recipe steps by uid %s: %v", uid, err)
		return []entity.RecipeStep{}, errors.WithStack(err)
	}

	return recipes, nil
}
