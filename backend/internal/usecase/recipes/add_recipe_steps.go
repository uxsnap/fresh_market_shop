package useCaseRecipes

import (
	"context"
	"log"

	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

func (uc *UseCaseRecipes) AddRecipeSteps(ctx context.Context, uid uuid.UUID, rSteps []entity.RecipeStep) error {
	log.Printf("ucRecipes.AddRecipeSteps")

	if err := validateRecipeSteps(rSteps); err != nil {
		log.Printf("failed to add recipe steps: %v", err)
		return errors.WithStack(err)
	}

	if err := uc.recipesRepository.AddRecipeSteps(ctx, uid, rSteps); err != nil {
		log.Printf("uc failed to add recipe steps: %v", err)
		return errors.WithStack(err)
	}

	return nil
}
