package useCaseOrders

import (
	"context"
	"log"
	"time"

	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

func (s *UseCaseOrders) CreateOrder(ctx context.Context, orderProducts entity.OrderProducts) (uuid.UUID, error) {
	log.Printf("ucOrders.CreateOrder")

	uuids := []uuid.UUID{}

	for _, u := range orderProducts.Products {
		uuids = append(uuids, u.Uid)
	}

	products, err := s.productsRepository.GetProductsWithExtra(
		ctx, uuid.Nil, 0, 0, time.Time{}, time.Time{}, 0, 0, true, false, uuids,
	)

	if err != nil {
		log.Printf("failed to fetch products for order: %v", err)
		return uuid.UUID{}, err
	}

	if err := validateOrderCreation(orderProducts, products); err != nil {
		log.Printf("failed to create order: %v", err)
		return uuid.UUID{}, err
	}

	order := entity.Order{
		Sum:       orderProducts.Sum,
		Uid:       uuid.NewV4(),
		CreatedAt: time.Now().UTC(),
	}

	if err := s.ordersRepository.CreateOrder(ctx, order); err != nil {
		log.Printf("failed to create order: %v", err)
		return uuid.UUID{}, errors.WithStack(err)
	}

	return order.Uid, nil
}
