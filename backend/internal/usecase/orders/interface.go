package useCaseOrders

import (
	"context"

	uuid "github.com/satori/go.uuid"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

type OrdersRepository interface {
	CreateOrder(ctx context.Context, order entity.Order) error
	UpdateOrder(ctx context.Context, order entity.Order) error
	GetOrder(ctx context.Context, qFilters entity.QueryFilters) (entity.Order, bool, error)
	GetOrderWithProducts(ctx context.Context, userUid uuid.UUID, qFilters entity.QueryFilters) ([]entity.OrderWithProducts, error)
}

type ProductsCountRepository interface {
	CheckIfAllItemsExist(ctx context.Context, productsCounts entity.ProductsCounts) error
	UpdateCount(ctx context.Context, productsCounts entity.ProductsCounts) error
}

type OrderProductsRepository interface {
	AddOrderProducts(ctx context.Context, orderProducts []entity.OrderProducts) error
}

type ProductsRepository interface {
	GetProductsWithExtra(ctx context.Context, qFilters entity.QueryFilters) ([]entity.ProductWithExtra, error)
	GetOrderedProductsSum(ctx context.Context, orderProducts []entity.OrderProducts) (int64, error)
}

type DeliveryService interface {
	GetDeliveryByOrderUid(ctx context.Context, orderUid uuid.UUID) (entity.Delivery, bool, error)
	GetDeliveryByUid(ctx context.Context, uid uuid.UUID) (entity.Delivery, bool, error)
	UpdateDelivery(ctx context.Context, delivery entity.Delivery) error
}

type PaymentsService interface {
	GetUserPaymentCardByUid(ctx context.Context, cardUid uuid.UUID) (entity.UserPaymentCard, bool, error)
	CreatePayment(ctx context.Context, payment entity.Payment) (uuid.UUID, error)
}
