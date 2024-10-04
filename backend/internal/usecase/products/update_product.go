package useCaseProducts

import (
	"context"
	"log"
	"time"

	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

func (uc *UseCaseProducts) UpdateProduct(ctx context.Context, product entity.Product) error {
	log.Printf("ucProducts.UpdateProduct: uid %s", product.Uid)

	if err := validateProduct(product); err != nil {
		log.Printf("failed to update product %s: %v", product.Uid, err)
		return errors.WithStack(err)
	}

	// TODO: подумать над объединением в транзакцию, чтобы транзакции обновления выполнялись последовательно (Serializable)
	savedProduct, isFound, err := uc.productsRepository.GetProductByUid(ctx, product.Uid)
	if err != nil {
		log.Printf("failed to update product: failed to get product %s: %v", product.Uid, err)
		return errors.WithStack(err)
	}

	if !isFound {
		log.Printf("failed to update product: product %s not found", product.Uid)
		return errors.New("product not found")
	}

	if !uuid.Equal(product.CategoryUid, savedProduct.CategoryUid) {
		_, isFound, err := uc.categoriesRepository.GetCategoryByUid(ctx, product.CategoryUid)
		if err != nil {
			log.Printf("failed to update product %s: %v", product.Uid, err)
			return errors.WithStack(err)
		}

		if !isFound {
			log.Printf("failed to update product: new category %s not found", product.CategoryUid)
			return errors.New("category not found")
		}
	}

	product.CreatedAt = savedProduct.CreatedAt
	product.UpdatedAt = time.Now().UTC()
	if err := uc.productsRepository.UpdateProduct(ctx, product); err != nil {
		log.Printf("failed to update product %s: %v", product.Uid, err)
		return errors.WithStack(err)
	}
	return nil
}
