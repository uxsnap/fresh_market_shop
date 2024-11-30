package useCaseOrders

import (
	"context"
	"log"

	"github.com/pkg/errors"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

func (uc *UseCaseOrders) GetOrder(ctx context.Context, qFilters entity.QueryFilters) (entity.Order, bool, error) {
	log.Printf("usecaseOrders.GetOrder: %s", qFilters.OrderUid)

	order, isFound, err := uc.ordersRepository.GetOrder(ctx, qFilters)
	if err != nil {
		log.Printf("failed to get order %s: %v", qFilters.OrderUid, err)
		return entity.Order{}, false, errors.WithStack(err)
	}

	return order, isFound, nil
}
