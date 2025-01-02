package useCaseProducts

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

const CATEGORY_KEY = "category"
const FILES_KEY = "file"

// продолжаем стараться записать в директорию
// файлы, которые упали, просто пропускаем
func (uc *UseCaseProducts) uploadProductFile(dir string, file *multipart.FileHeader) error {
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

func (uc *UseCaseProducts) UploadProductPhotos(ctx context.Context, productUid uuid.UUID, form *multipart.Form) error {
	log.Printf("ucProducts.UploadProductPhotos: uid %s", productUid)

	categoryUid, err := uuid.FromString(form.Value[CATEGORY_KEY][0])
	if err != nil {
		log.Printf("failed to get category name %s", productUid)
		return errors.New("no category name")
	}

	categoryName, ok := utils.MapCategoryIdToCategoryName[categoryUid]
	if !ok {
		log.Printf("failed to get category name %s", productUid)
		return errors.New("no category name")
	}

	uploadPath, _ := os.Getwd()

	serverDir := filepath.Join("assets", "imgs", categoryName)
	productDir := filepath.Join(uploadPath, serverDir)
	if _, err := os.Stat(productDir); err != nil {
		log.Printf("failed to get category name %s %v", productUid, err)
		return errors.New("no category name")
	}

	files, ok := form.File[FILES_KEY]
	if !ok {
		log.Printf("failed to get files %s", productUid)
		return errors.New("no category files")
	}

	uploadedImgPaths := []string{}

	for _, file := range files {
		if err := uc.uploadProductFile(productDir, file); err == nil {
			uploadedImgPaths = append(uploadedImgPaths, filepath.Join(serverDir, file.Filename))
		}
	}

	if err := uc.productsRepository.UpdateProductPhotos(ctx, productUid, uploadedImgPaths); err != nil {
		log.Printf("failed to update image paths %s, %v", productUid, err)
		return errors.New("no category files")
	}

	return nil
}
