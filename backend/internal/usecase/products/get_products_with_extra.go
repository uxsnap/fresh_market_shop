package useCaseProducts

import (
	"context"
	"log"

	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

func (uc *UseCaseProducts) GetProductsWithExtra(ctx context.Context, qFilters entity.QueryFilters) ([]entity.ProductWithExtra, error) {
	log.Printf("ucProducts.GetProductsWithExtra")

	categoryUid := qFilters.CategoryUid

	if !uuid.Equal(categoryUid, uuid.UUID{}) {
		_, categoryFound, err := uc.categoriesRepository.GetCategoryByUid(ctx, categoryUid)
		if err != nil {
			log.Printf("failed to get category %s: %v", categoryUid, err)
		}

		if !categoryFound {
			log.Printf("category %s not found", categoryUid)
			return nil, errors.New("category not found")
		}
	}

	products, err := uc.productsRepository.GetProductsWithExtra(ctx, qFilters)
	if err != nil {
		log.Printf("failed to get products: %v", err)
		return nil, errors.WithStack(err)
	}
	return products, nil
}
