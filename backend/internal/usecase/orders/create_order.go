package useCaseOrders

import (
	"context"
	"log"

	uuid "github.com/satori/go.uuid"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
	errorWrapper "github.com/uxsnap/fresh_market_shop/backend/internal/error_wrapper"
)

func (s *UseCaseOrders) CreateOrder(ctx context.Context, orderProducts entity.OrderProducts) (uuid.UUID, *errorWrapper.Error) {
	log.Printf("ucOrders.CreateOrder")

	if err := s.productsRepository.CheckIfAllItemsExist(ctx, orderProducts); err != nil {
		log.Printf("failed to validate order creation: %v", err)
		return uuid.UUID{}, err
	}

	order := entity.Order{
		Uid: uuid.NewV4(),
	}

	if err := s.ordersRepository.CreateOrder(ctx, order); err != nil {
		log.Printf("failed to create order: %v", err)
		return uuid.UUID{}, err
	}

	if err := s.productsRepository.UpdateCount(ctx, orderProducts); err != nil {
		log.Printf("failed to update products count: %v", err)
		return uuid.UUID{}, err
	}

	return order.Uid, nil
}
