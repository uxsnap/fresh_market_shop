package useCaseOrders

import (
	"context"
	"log"

	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

func (s *UseCaseOrders) CreateOrder(ctx context.Context, userUid uuid.UUID, productsCounts entity.ProductsCounts) (uuid.UUID, error) {
	log.Printf("ucOrders.CreateOrder: user uid %s", userUid)

	order := entity.Order{
		UserUid: userUid,
		Uid:     uuid.NewV4(),
	}

	orderProducts := make([]entity.OrderProducts, len(productsCounts.Products))
	for ind, val := range productsCounts.Products {
		orderProducts[ind] = entity.OrderProducts{
			OrderUid:   order.Uid,
			ProductUid: val.ProductUid,
			Count:      val.Count,
		}
	}

	if err := s.txManager.NewPgTransaction().Execute(ctx, func(ctx context.Context) error {

		if err := s.productsCountRepository.CheckIfAllItemsExist(ctx, productsCounts); err != nil {
			log.Printf("failed to validate order creation: %v", err)
			return err
		}

		if err := s.productsCountRepository.UpdateCount(ctx, productsCounts); err != nil {
			log.Printf("failed to update products count: %v", err)
			return err
		}

		if err := s.ordersRepository.CreateOrder(ctx, order); err != nil {
			log.Printf("failed to create order: %v", err)
			return err
		}

		if err := s.orderProductsRepository.AddOrderProducts(ctx, orderProducts); err != nil {
			log.Printf("failed to create order: %v", err)
			return err
		}
		return nil
	}); err != nil {
		return uuid.UUID{}, errors.WithStack(err)
	}

	return order.Uid, nil
}
