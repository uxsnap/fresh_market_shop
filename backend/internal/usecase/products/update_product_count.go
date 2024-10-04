package useCaseProducts

import (
	"context"
	"log"

	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
)

func (uc *UseCaseProducts) UpdateProductCount(ctx context.Context, productUid uuid.UUID, stockQuantity int64) error {
	log.Printf("ucProducts.UpdateProductCount: productUid: %s count: %d", productUid, stockQuantity)

	if err := uc.txManager.NewPgTransaction().Execute(ctx, func(ctx context.Context) error {

		_, isFound, err := uc.productsRepository.GetProductCount(ctx, productUid)
		if err != nil {
			return err
		}
		if !isFound {
			return uc.productsRepository.CreateProductCount(ctx, productUid, stockQuantity)
		}
		return uc.productsRepository.UpdateProductCount(ctx, productUid, stockQuantity)
	}); err != nil {
		log.Printf("failed to update product(uid: %s) count: %v", productUid, stockQuantity)
		return errors.WithStack(err)
	}

	return nil
}
