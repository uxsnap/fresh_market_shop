package useCaseRecipes

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"github.com/uxsnap/fresh_market_shop/backend/internal/utils"
)

func (uc *UseCaseRecipes) deleteRecipeFile(dir string, name string) error {
	if !utils.IsImageExtensionAllowed(name) {
		log.Println("error file extension")
		return errors.New("error file extension")
	}

	path := filepath.Join(dir, name)

	err := os.Remove(path)
	if err != nil {
		log.Println("error deleting file", err)
		return errors.New("error deleting file")
	}

	return nil
}

func (uc *UseCaseRecipes) DeleteRecipePhotos(ctx context.Context, recipeUid uuid.UUID, photoNames ...string) error {
	log.Printf("ucRecipes.DeleteRecipePhotos: product uid: %s photos: %v", recipeUid, photoNames)

	if uuid.Equal(recipeUid, uuid.UUID{}) {
		log.Printf("failed to delete product photos: empty product uid")
		return errors.New("failed to delete product photos: empty product uid")
	}

	if len(photoNames) == 0 {
		fmt.Printf("failed to delete product %s photos: empty photos uids", recipeUid)
		return errors.New("empty product photos uids")
	}

	deletePath, _ := os.Getwd()

	recipeDir := filepath.Join(deletePath, "assets", "recipes", recipeUid.String())
	if _, err := os.Stat(recipeDir); err != nil {
		log.Printf("failed to get folder %s %v", recipeUid, err)
		return errors.New("no folder")
	}

	for _, name := range photoNames {
		uc.deleteRecipeFile(recipeDir, name)
	}

	return nil
}
