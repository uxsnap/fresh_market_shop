package useCaseRecipes

import (
	"context"
	"log"

	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
)

func (uc *UseCaseRecipes) DeleteRecipe(ctx context.Context, uid uuid.UUID) error {
	log.Printf("ucRecipes.DeleteRecipe: uid %s", uid)

	if err := uc.recipesRepository.DeleteRecipe(ctx, uid); err != nil {
		log.Printf("failed to delete recipe %s: %v", uid, err)
		return errors.WithStack(err)
	}

	return nil
}
