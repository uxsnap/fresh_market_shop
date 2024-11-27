package useCaseOrders

import (
	"context"
	"log"

	uuid "github.com/satori/go.uuid"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

func (uc *UseCaseOrders) GetOrderHistory(ctx context.Context, userUid uuid.UUID, qFilters entity.QueryFilters) ([]entity.OrderWithProducts, error) {
	log.Printf("ucOrders.GetOrderHistory: user uid %s", userUid)

	orderWithProducts, err := uc.ordersRepository.GetOrderWithProducts(ctx, userUid, qFilters)

	if err != nil {
		log.Printf("failed to get order history %s: %v", userUid, err)
		return []entity.OrderWithProducts{}, err
	}

	return orderWithProducts, nil
}
