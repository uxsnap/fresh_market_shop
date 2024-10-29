package useCaseProducts

import (
	"context"
	"log"

	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

func (uc *UseCaseProducts) GetProductByUid(ctx context.Context, productUid uuid.UUID) (entity.Product, bool, error) {
	log.Printf("ucProducts.GetProductByUid: uid: %s", productUid)

	product, isFound, err := uc.productsRepository.GetProductByUid(ctx, productUid)
	return product, isFound, errors.WithStack(err)
}

func (uc *UseCaseProducts) GetProducts(ctx context.Context, qFilters entity.QueryFilters) ([]entity.Product, error) {
	log.Printf("ucProducts.GetProducts")

	if !uuid.Equal(qFilters.CategoryUid, uuid.UUID{}) {
		_, categoryFound, err := uc.categoriesRepository.GetCategoryByUid(ctx, qFilters.CategoryUid)
		if err != nil {
			log.Printf("failed to get category %s: %v", qFilters.CategoryUid, err)
		}

		if !categoryFound {
			log.Printf("category %s not found", qFilters.CategoryUid)
			return nil, errors.New("category not found")
		}
	}

	products, err := uc.productsRepository.GetProducts(ctx, qFilters)
	if err != nil {
		log.Printf("failed to get products: %v", err)
		return nil, errors.WithStack(err)
	}
	return products, nil
}
