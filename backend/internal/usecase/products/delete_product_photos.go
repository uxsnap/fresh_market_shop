package useCaseProducts

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
)

func (uc *UseCaseProducts) DeleteProductPhotos(ctx context.Context, productUid uuid.UUID, photosUids ...uuid.UUID) error {
	log.Printf("ucProducts.DeleteProductPhotos: product uid: %s photos: %v", productUid, photosUids)

	if uuid.Equal(productUid, uuid.UUID{}) {
		log.Printf("failed to delete product photos: empty product uid")
		return errors.New("failed to delete product photos: empty product uid")
	}

	if len(photosUids) == 0 {
		fmt.Printf("failed to delete product %s photos: empty photos uids", productUid)
		return errors.New("empty product photos uids")
	}

	photos, err := uc.productsRepository.GetProductPhotos(ctx, productUid, photosUids...)
	if err != nil {
		fmt.Printf("failed to delete product %s photos: %v", productUid, err)
		return errors.WithStack(err)
	}

	if err := uc.productsRepository.DeleteProductPhotos(ctx, productUid, photosUids...); err != nil {
		log.Printf("failed to delete product %s photos: %v", productUid, err)
		return errors.WithStack(err)
	}

	//TODO: подумать над этим местом
	for _, photo := range photos {
		_ = uc.deleteProductFile(photo.FilePath)
	}
	return nil
}

func (uc *UseCaseProducts) deleteProductFile(filePath string) error {
	err := os.Remove(filePath)
	if err != nil {
		log.Printf("WARN: error removing file %s: %v", filePath, err)
		return errors.Errorf("error removing file %s: %v", filePath, err)
	}
	return nil
}
