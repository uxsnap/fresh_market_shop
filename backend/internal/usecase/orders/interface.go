package useCaseOrders

import (
	"context"

	uuid "github.com/satori/go.uuid"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

type OrdersRepository interface {
	CreateOrder(ctx context.Context, order entity.Order) error
}

type ProductsRepository interface {
	CheckIfAllItemsExist(ctx context.Context, uuids []uuid.UUID) error
}
