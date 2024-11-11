package useCaseDelivery

import (
	"context"

	uuid "github.com/satori/go.uuid"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

type DeliveryRepository interface {
	CreateDelivery(ctx context.Context, delivery entity.Delivery) error
	GetDeliveryByOrderUid(ctx context.Context, orderUid uuid.UUID) (entity.Delivery, bool, error)
	GetDeliveryByUid(ctx context.Context, uid uuid.UUID) (entity.Delivery, bool, error)
	UpdateDelivery(ctx context.Context, delivery entity.Delivery) error
	GetDeliveryHistoryByUser(ctx context.Context, userUid uuid.UUID) ([]entity.Delivery, error)
}
