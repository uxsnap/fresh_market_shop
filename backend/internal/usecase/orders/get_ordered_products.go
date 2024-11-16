package useCaseOrders

import (
	"context"
	"log"

	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

func (uc *UseCaseOrders) GetOrderedProducts(ctx context.Context, qFilters entity.QueryFilters) ([]entity.ProductWithExtra, error) {
	log.Printf("ucOrders.GetOrderedProducts %v", qFilters.OrderUid)

	products, err := uc.productsRepository.GetProductsWithExtra(ctx, qFilters)
	if err != nil {
		log.Printf("failed to get order products: %v", err)
		return nil, err
	}
	return products, nil
}
