package useCaseProducts

import (
	"context"
	"log"

	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
)

func (uc *UseCaseProducts) DeleteCategory(ctx context.Context, uid uuid.UUID) error {
	log.Printf("ucProducts.DeleteCategory: uid %s", uid)

	if err := uc.categoriesRepository.DeleteCategory(ctx, uid); err != nil {
		log.Printf("failed to delete category %s: %v", uid, err)
		return errors.WithStack(err)
	}
	return nil
}
