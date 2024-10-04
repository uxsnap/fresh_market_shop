package useCaseProducts

import (
	"context"
	"log"

	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
)

func (uc *UseCaseProducts) DeleteProduct(ctx context.Context, productUid uuid.UUID) error {
	log.Printf("ucProducts.DeleteProduct: uid %s", productUid)

	if err := uc.productsRepository.DeleteProduct(ctx, productUid); err != nil {
		log.Printf("failed to delete product %s: %v", productUid, err)
		return errors.WithStack(err)
	}
	return nil
}
