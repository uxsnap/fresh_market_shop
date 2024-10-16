package useCaseRecipes

import (
	"context"
	"log"

	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

func (uc *UseCaseRecipes) GetRecipeByUid(ctx context.Context, uid uuid.UUID) (entity.Recipe, bool, error) {
	log.Printf("ucRecipes.GetRecipeByUid: uid %s", uid)

	recipe, isFound, err := uc.recipesRepository.GetRecipeByUid(ctx, uid)
	if err != nil {
		log.Printf("failed to get recipe by uid %s: %v", uid, err)
		return entity.Recipe{}, false, errors.WithStack(err)
	}
	if !isFound {
		return entity.Recipe{}, false, nil
	}

	return recipe, true, nil
}
