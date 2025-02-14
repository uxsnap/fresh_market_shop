package useCaseRecipes

import (
	"context"
	"log"

	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
)

func (uc *UseCaseRecipes) DeleteRecipeStep(ctx context.Context, uid uuid.UUID, step int) error {
	log.Printf("ucRecipes.DeleteRecipeStep: step %v", step)

	if err := uc.recipesRepository.DeleteRecipeStep(ctx, uid, step); err != nil {
		log.Printf("failed to delete recipe step %v: %v", step, err)
		return errors.WithStack(err)
	}

	return nil
}
