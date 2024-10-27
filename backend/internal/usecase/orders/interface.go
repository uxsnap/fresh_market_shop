package useCaseOrders

import (
	"context"

	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
	errorWrapper "github.com/uxsnap/fresh_market_shop/backend/internal/error_wrapper"
)

type OrdersRepository interface {
	CreateOrder(ctx context.Context, order entity.Order) *errorWrapper.Error
}

type ProductsCountRepository interface {
	CheckIfAllItemsExist(ctx context.Context, productsCounts entity.ProductsCounts) *errorWrapper.Error
	UpdateCount(ctx context.Context, productsCounts entity.ProductsCounts) *errorWrapper.Error
}

type OrderProductsRepository interface {
	AddOrderProducts(ctx context.Context, orderProducts []entity.OrderProducts) *errorWrapper.Error
}
