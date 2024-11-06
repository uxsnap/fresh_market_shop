package useCaseOrders

import (
	"context"
	"log"

	uuid "github.com/satori/go.uuid"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

func (s *UseCaseOrders) CreateOrder(ctx context.Context, productsCounts entity.ProductsCounts) (uuid.UUID, error) {
	log.Printf("ucOrders.CreateOrder")

	if err := s.productsCountRepository.CheckIfAllItemsExist(ctx, productsCounts); err != nil {
		log.Printf("failed to validate order creation: %v", err)
		return uuid.UUID{}, err
	}

	if err := s.productsCountRepository.UpdateCount(ctx, productsCounts); err != nil {
		log.Printf("failed to update products count: %v", err)
		return uuid.UUID{}, err
	}

	order := entity.Order{
		Uid: uuid.NewV4(),
	}

	if err := s.ordersRepository.CreateOrder(ctx, order); err != nil {
		log.Printf("failed to create order: %v", err)
		return uuid.UUID{}, err
	}

	orderProducts := make([]entity.OrderProducts, len(productsCounts.Products))

	for ind, val := range productsCounts.Products {
		orderProducts[ind] = entity.OrderProducts{
			OrderUid:   order.Uid,
			ProductUid: val.ProductUid,
			Count:      val.Count,
		}
	}

	if err := s.orderProductsRepository.AddOrderProducts(ctx, orderProducts); err != nil {
		log.Printf("failed to create order: %v", err)
		return uuid.UUID{}, err
	}

	return order.Uid, nil
}
