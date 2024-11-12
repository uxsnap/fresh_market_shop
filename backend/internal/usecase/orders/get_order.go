package useCaseOrders

import (
	"context"
	"log"

	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

func (uc *UseCaseOrders) GetOrder(ctx context.Context, orderUid uuid.UUID) (entity.Order, bool, error) {
	log.Printf("usecaseOrders.GetOrder: %s", orderUid)

	order, isFound, err := uc.ordersRepository.GetOrderByUid(ctx, orderUid)
	if err != nil {
		log.Printf("failed to get order %s: %v", orderUid, err)
		return entity.Order{}, false, errors.WithStack(err)
	}

	return order, isFound, nil
}
