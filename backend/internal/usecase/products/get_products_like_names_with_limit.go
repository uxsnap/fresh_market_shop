package useCaseProducts

import (
	"context"
	"log"

	"github.com/pkg/errors"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

func (uc *UseCaseProducts) GetProductsLikeNamesWithLimitOnEach(ctx context.Context, names []string, qFilters entity.QueryFilters) ([]entity.Product, error) {
	log.Printf("ucProducts.GetProductsLikeNamesWithLimitOnEach: limit on each %d names %v", qFilters.Limit, names)

	if len(names) == 0 {
		return []entity.Product{}, nil
	}

	products, err := uc.productsRepository.GetProductsLikeNamesWithLimitOnEach(ctx, names, qFilters)
	if err != nil {
		log.Printf("failed to get GetProductsLikeNamesWithLimitOnEach: %v", err)
		return nil, errors.WithStack(err)
	}

	return products, nil
}

func (uc *UseCaseProducts) GetProductsLikeNamesWithLimitOnEachWithExtra(ctx context.Context, names []string, qFilters entity.QueryFilters) ([]entity.ProductWithExtra, error) {
	log.Printf("ucProducts.GetProductsLikeNamesWithLimitOnEachWithExtra: limit on each %d names %v", qFilters.Limit, names)

	if len(names) == 0 {
		return []entity.ProductWithExtra{}, nil
	}

	products, err := uc.productsRepository.GetProductsLikeNamesWithLimitOnEachWithExtra(ctx, names, qFilters)
	if err != nil {
		log.Printf("failed to get GetProductsLikeNamesWithLimitOnEachWithExtra: %v", err)
		return nil, errors.WithStack(err)
	}

	return products, nil
}
