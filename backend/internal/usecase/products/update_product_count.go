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

func (uc *UseCaseProducts) IncrementProductCount(ctx context.Context, productUid uuid.UUID, incValue int64) error {
	log.Printf("ucProducts.IncrementProductCount product uid: %s, incValue: %d", productUid, incValue)

	if err := uc.changeProductCount(ctx, productUid, incValue, true); err != nil {
		log.Printf("failed to increment product(uid: %s) count:%v", productUid, err)
		return errors.WithStack(err)
	}
	return nil
}

func (uc *UseCaseProducts) DecrementProductCount(ctx context.Context, productUid uuid.UUID, decValue int64) error {
	log.Printf("ucProducts.DecrementProductCount product uid: %s, incValue: %d", productUid, decValue)

	if err := uc.changeProductCount(ctx, productUid, decValue, false); err != nil {
		log.Printf("failed to decrement product(uid: %s) count:%v", productUid, err)
		return errors.WithStack(err)
	}
	return nil
}

func (uc *UseCaseProducts) changeProductCount(ctx context.Context, productUid uuid.UUID, value int64, inc bool) error {
	return uc.txManager.NewPgTransaction().Execute(ctx, func(ctx context.Context) error {
		pc, isFound, err := uc.productsRepository.GetProductCount(ctx, productUid)
		if err != nil {
			return err
		}
		if !isFound {
			return errors.Errorf("product count with product uid %s not found", productUid)
		}

		var newCount int64
		if inc {
			newCount = pc + value
		} else {
			newCount = pc - value
			if newCount < 0 {
				return errors.New("new product count < 0")
			}
		}

		return uc.productsRepository.UpdateProductCount(ctx, productUid, newCount)
	})
}
