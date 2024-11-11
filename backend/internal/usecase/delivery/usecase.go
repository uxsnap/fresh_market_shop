package useCaseDelivery

import (
	"github.com/uxsnap/fresh_market_shop/backend/internal/manager/transaction"
)

type UseCaseDelivery struct {
	deliveryRepository DeliveryRepository
	txManager          *transaction.Manager
}

func New(
	deliveryRepository DeliveryRepository,
	txManager *transaction.Manager,
) *UseCaseDelivery {
	return &UseCaseDelivery{
		deliveryRepository: deliveryRepository,
		txManager:          txManager,
	}
}

const (
	fromLongitude = 54.711712
	fromLatitude  = 20.579137
)
