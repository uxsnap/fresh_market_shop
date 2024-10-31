package useCaseRecipes

import (
	"context"
	"log"

	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

func (uc *UseCaseRecipes) GetRecipesProducts(ctx context.Context, uid uuid.UUID) ([]entity.ProductWithExtra, error) {
	log.Printf("ucRecipes.GetRecipeByUid: uid %s", uid)

	recipes, err := uc.productsRepository.GetProductsWithExtra(ctx, entity.QueryFilters{RecipeUid: uid})

	if err != nil {
		log.Printf("failed to get recipe by uid %s: %v", uid, err)
		return []entity.ProductWithExtra{}, errors.WithStack(err)
	}

	return recipes, nil
}
