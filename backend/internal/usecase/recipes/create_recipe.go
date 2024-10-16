package useCaseRecipes

import (
	"context"
	"log"
	"time"

	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

func (uc *UseCaseRecipes) CreateRecipe(ctx context.Context, recipe entity.Recipe) (uuid.UUID, error) {
	log.Printf("ucRecipes.CreateRecipe: name %s", recipe.Name)

	if err := validateRecipe(recipe); err != nil {
		log.Printf("failed to create recipe: %v", err)
		return uuid.UUID{}, errors.WithStack(err)
	}

	recipe.Uid = uuid.NewV4()
	recipe.CreatedAt = time.Now().UTC()

	if err := uc.recipesRepository.CreateRecipe(ctx, recipe); err != nil {
		log.Printf("failed to create recipe: %v", err)
		return uuid.UUID{}, errors.WithStack(err)
	}

	return recipe.Uid, nil
}
