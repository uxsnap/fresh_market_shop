package useCaseOrders

import (
	"context"
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

type OrdersRepository interface {
	CreateOrder(ctx context.Context, order entity.Order) error
}

type ProductsRepository interface {
	GetProductsWithExtra(ctx context.Context, categoryUid uuid.UUID, ccalMin int64, ccalMax int64, createdBefore time.Time, createdAfter time.Time, limit uint64, offset uint64, withCounts bool, withPhotos bool, uuids []uuid.UUID) ([]entity.ProductWithExtra, error)
}
