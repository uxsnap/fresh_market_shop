package useCaseProducts

import (
	"context"
	"log"

	"github.com/pkg/errors"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

func (uc *UseCaseProducts) GetProductsByNameLike(ctx context.Context, name string, limit uint64, offset uint64) ([]entity.Product, error) {
	log.Printf("ucProducts.GetProductsByNameLike: name %s", name)

	if len(name) == 0 {
		return nil, nil
	}

	products, err := uc.productsRepository.GetProductsByNameLike(ctx, name, limit, offset)
	if err != nil {
		log.Printf("failed to get products by name like %s: %v", name, err)
		return nil, errors.WithStack(err)
	}

	return products, nil
}

func (uc *UseCaseProducts) GetProductsByNameLikeWithCounts(ctx context.Context, name string, limit uint64, offset uint64) ([]entity.ProductWithStockQuantity, error) {
	log.Printf("ucProducts.GetProductsByNameLikeWithCounts: name %s", name)

	if len(name) == 0 {
		return nil, nil
	}

	products, err := uc.productsRepository.GetProductsByNameLikeWithCounts(ctx, name, limit, offset)
	if err != nil {
		log.Printf("failed to get products by name like with counts %s: %v", name, err)
		return nil, errors.WithStack(err)
	}

	return products, nil
}
