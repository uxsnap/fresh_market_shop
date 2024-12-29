package useCaseProducts

import (
	"context"
	"log"

	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
)

func (uc *UseCaseProducts) ReviveProduct(ctx context.Context, productUid uuid.UUID) error {
	log.Printf("ucProducts.ReviveProduct: uid %s", productUid)

	if err := uc.productsRepository.ReviveProduct(ctx, productUid); err != nil {
		log.Printf("failed to revive product %s: %v", productUid, err)
		return errors.WithStack(err)
	}
	return nil
}
