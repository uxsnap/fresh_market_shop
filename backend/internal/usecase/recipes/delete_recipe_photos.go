package useCaseRecipes

import (
	"context"
	"fmt"
	"log"

	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
)

func (uc *UseCaseRecipes) DeleteRecipePhotos(ctx context.Context, recipeUid uuid.UUID, photosUids ...uuid.UUID) error {
	log.Printf("ucRecipes.DeleteRecipePhotos: product uid: %s photos: %v", recipeUid, photosUids)

	if uuid.Equal(recipeUid, uuid.UUID{}) {
		log.Printf("failed to delete product photos: empty product uid")
		return errors.New("failed to delete product photos: empty product uid")
	}

	if len(photosUids) == 0 {
		fmt.Printf("failed to delete product %s photos: empty photos uids", recipeUid)
		return errors.New("empty product photos uids")
	}

	if err := uc.recipesRepository.DeleteRecipePhotos(ctx, recipeUid, photosUids...); err != nil {
		log.Printf("failed to delete product %s photos: %v", recipeUid, err)
		return errors.WithStack(err)
	}

	return nil
}
