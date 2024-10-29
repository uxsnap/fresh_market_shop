package useCaseProducts

import (
	"context"
	"log"

	"github.com/pkg/errors"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

func (uc *UseCaseProducts) GetProductsByNameLike(ctx context.Context, name string, qFilters entity.QueryFilters) ([]entity.Product, error) {
	log.Printf("ucProducts.GetProductsByNameLike: name %s", name)

	if len(name) == 0 {
		return nil, nil
	}

	products, err := uc.productsRepository.GetProductsByNameLike(ctx, name, qFilters)
	if err != nil {
		log.Printf("failed to get products by name like %s: %v", name, err)
		return nil, errors.WithStack(err)
	}

	return products, nil
}

func (uc *UseCaseProducts) GetProductsByNameLikeWithExtra(ctx context.Context, name string, qFilters entity.QueryFilters) ([]entity.ProductWithExtra, error) {
	log.Printf("ucProducts.GetProductsByNameLikeWithExtra: name %s", name)

	if len(name) == 0 {
		return nil, nil
	}

	products, err := uc.productsRepository.GetProductsByNameLikeWithExtra(ctx, name, qFilters)
	if err != nil {
		log.Printf("failed to get products by name like with extra %s: %v", name, err)
		return nil, errors.WithStack(err)
	}

	return products, nil
}
