package useCaseProducts

import (
	"context"
	"log"

	"github.com/pkg/errors"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

func (uc *UseCaseProducts) GetProductsLikeNamesWithLimitOnEach(ctx context.Context, names []string, limit uint64, offset uint64) ([]entity.Product, error) {
	log.Printf("ucProducts.GetProductsLikeNamesWithLimitOnEach: limit on each %d names %v", limit, names)

	if len(names) == 0 {
		return []entity.Product{}, nil
	}

	products, err := uc.productsRepository.GetProductsLikeNamesWithLimitOnEach(ctx, names, limit, offset)
	if err != nil {
		log.Printf("failed to get GetProductsLikeNamesWithLimitOnEach: %v", err)
		return nil, errors.WithStack(err)
	}

	return products, nil
}

func (uc *UseCaseProducts) GetProductsLikeNamesWithLimitOnEachWithCounts(ctx context.Context, names []string, limit uint64, offset uint64) ([]entity.ProductWithStockQuantity, error) {
	log.Printf("ucProducts.GetProductsLikeNamesWithLimitOnEachWithCounts: limit on each %d names %v", limit, names)

	if len(names) == 0 {
		return []entity.ProductWithStockQuantity{}, nil
	}

	products, err := uc.productsRepository.GetProductsLikeNamesWithLimitOnEachWithCounts(ctx, names, limit, offset)
	if err != nil {
		log.Printf("failed to get GetProductsLikeNamesWithLimitOnEachWithCounts: %v", err)
		return nil, errors.WithStack(err)
	}

	return products, nil
}