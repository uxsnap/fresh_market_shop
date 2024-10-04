package useCaseProducts

import (
	"context"
	"log"

	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
)

func (uc *UseCaseProducts) GetProductCount(ctx context.Context, productUid uuid.UUID) (int64, bool, error) {
	log.Printf("ucProduct.GetProductCount: product uid: %s", productUid)

	count, isFound, err := uc.productsRepository.GetProductCount(ctx, productUid)
	if err != nil {
		log.Printf("failed to get count of product %s: %v", productUid, err)
		return count, isFound, errors.WithStack(err)
	}
	return count, isFound, nil
}
