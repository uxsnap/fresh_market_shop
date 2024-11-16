package useCaseDelivery

import (
	"github.com/uxsnap/fresh_market_shop/backend/internal/manager/transaction"
)

type UseCaseDelivery struct {
	deliveryRepository DeliveryRepository
	usersService       UsersService
	txManager          *transaction.Manager
}

func New(
	deliveryRepository DeliveryRepository,
	usersService UsersService,
	txManager *transaction.Manager,
) *UseCaseDelivery {
	return &UseCaseDelivery{
		deliveryRepository: deliveryRepository,
		usersService:       usersService,
		txManager:          txManager,
	}
}

const (
	fromLongitude = 54.711712
	fromLatitude  = 20.579137
)
