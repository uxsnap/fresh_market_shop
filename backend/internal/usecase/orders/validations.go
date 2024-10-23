package useCaseOrders

import (
	"errors"

	uuid "github.com/satori/go.uuid"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

func validateOrderCreation(order entity.OrderProducts, products []entity.ProductWithExtra) error {
	m := make(map[uuid.UUID]int64)

	if len(order.Products) != len(products) {
		return errors.New("неправильно созданный заказ")
	}

	for _, op := range order.Products {
		m[op.Uid] = op.Count
	}

	for _, p := range products {
		if p.StockQuantity < m[p.Uid] {
			return errors.New("недостаточно товаров на складе для совершения заказа")
		}
	}

	return nil
}
