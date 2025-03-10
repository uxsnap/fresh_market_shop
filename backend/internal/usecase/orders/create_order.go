package useCaseOrders

import (
	"context"
	"log"

	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

func (uc *UseCaseOrders) CreateOrder(ctx context.Context, userUid uuid.UUID, productsCounts entity.ProductsCounts) (uuid.UUID, error) {
	log.Printf("ucOrders.CreateOrder: user uid %s", userUid)

	order := entity.Order{
		UserUid: userUid,
		Status:  entity.OrderStatusNew,
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

	if err := uc.txManager.NewPgTransaction().Execute(ctx, func(ctx context.Context) error {

		if err := uc.productsCountRepository.CheckIfAllItemsExist(ctx, productsCounts); err != nil {
			log.Printf("failed to validate order creation: %v", err)
			return err
		}

		if err := uc.productsCountRepository.UpdateCount(ctx, productsCounts); err != nil {
			log.Printf("failed to update products count: %v", err)
			return err
		}

		if err := uc.orderProductsRepository.AddOrderProducts(ctx, orderProducts); err != nil {
			log.Printf("failed to create order: %v", err)
			return err
		}

		sum, err := uc.productsRepository.GetOrderedProductsSum(ctx, orderProducts)
		if err != nil {
			log.Printf("failed to get order sum: %v", err)
			return err
		}
		if sum < 0 {
			log.Printf("failed to get order sum: sum < 0")
			return errors.New("sum < 0")
		}
		order.Sum = sum

		if err := uc.ordersRepository.CreateOrder(ctx, order); err != nil {
			log.Printf("failed to create order: %v", err)
			return err
		}

		return nil
	}); err != nil {
		return uuid.UUID{}, errors.WithStack(err)
	}

	return order.Uid, nil
}
