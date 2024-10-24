package useCaseOrders

import (
	"context"
	"log"

	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

func (s *UseCaseOrders) CreateOrder(ctx context.Context, orderProducts entity.OrderProducts) (uuid.UUID, error) {
	log.Printf("ucOrders.CreateOrder")

	uuids := make([]uuid.UUID, len(orderProducts.Products))

	for ind, v := range orderProducts.Products {
		uuids[ind] = v.Uid
	}

	if err := s.productsRepository.CheckIfAllItemsExist(ctx, uuids); err != nil {
		log.Printf("failed to validate order creation: %v", err)
		return uuid.UUID{}, errors.WithStack(err)
	}

	order := entity.Order{
		Sum: orderProducts.Sum,
		Uid: uuid.NewV4(),
	}

	if err := s.ordersRepository.CreateOrder(ctx, order); err != nil {
		log.Printf("failed to create order: %v", err)
		return uuid.UUID{}, errors.WithStack(err)
	}

	return order.Uid, nil
}
