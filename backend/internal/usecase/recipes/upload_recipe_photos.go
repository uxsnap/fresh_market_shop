package useCaseRecipes

import (
	"context"
	"io"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"github.com/uxsnap/fresh_market_shop/backend/internal/utils"
)

const FILES_KEY = "file"

// продолжаем стараться записать в директорию
// файлы, которые упали, просто пропускаем

func (uc *UseCaseRecipes) uploadRecipeFile(dir string, file *multipart.FileHeader) error {
	curFile, err := file.Open()
	if err != nil {
		log.Println("error opening file", err)
		return errors.New("error opening file")
	}

	defer curFile.Close()

	if !utils.IsImageExtensionAllowed(file.Filename) {
		log.Println("error file extension", err)
		return errors.New("error file extension")
	}

	path := filepath.Join(dir, file.Filename)

	dst, err := os.Create(path)
	if err != nil {
		log.Println("error creating file", err)
		return errors.New("error creating file")
	}

	defer dst.Close()

	if _, err := io.Copy(dst, curFile); err != nil {
		log.Println("error copying file", err)
		return errors.New("error copying file")
	}

	return nil
}

func (uc *UseCaseRecipes) UploadRecipePhotos(ctx context.Context, recipeUid uuid.UUID, form *multipart.Form) error {
	log.Printf("ucRecipes.UploadRecipePhotos: uid %s", recipeUid)

	uploadPath, _ := os.Getwd()

	recipeDir := filepath.Join(uploadPath, "assets", "recipes", recipeUid.String())
	if err := os.MkdirAll(recipeDir, os.ModePerm); err != nil {
		log.Printf("failed to get folder %s %v", recipeUid, err)
		return errors.New("no folder")
	}

	files, ok := form.File[FILES_KEY]
	if !ok {
		log.Printf("failed to get files %s", recipeUid)
		return errors.New("no category files")
	}

	for _, file := range files {
		uc.uploadRecipeFile(recipeDir, file)
	}

	return nil
}
