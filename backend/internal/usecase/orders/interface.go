package useCaseOrders

import (
	"context"

	uuid "github.com/satori/go.uuid"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

type OrdersRepository interface {
	CreateOrder(ctx context.Context, order entity.Order) error
	GetOrderByUid(ctx context.Context, uid uuid.UUID) (entity.Order, bool, error)
}

type ProductsCountRepository interface {
	CheckIfAllItemsExist(ctx context.Context, productsCounts entity.ProductsCounts) error
	UpdateCount(ctx context.Context, productsCounts entity.ProductsCounts) error
}

type OrderProductsRepository interface {
	AddOrderProducts(ctx context.Context, orderProducts []entity.OrderProducts) error
}
