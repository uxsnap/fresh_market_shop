package useCasePayments

import "github.com/uxsnap/fresh_market_shop/backend/internal/manager/transaction"

type UseCasePayments struct {
	paymentsRepository PaymentsRepository
	usersService       UsersService
	ordersService      OrdersService
	txManager          *transaction.Manager
}

func New(
	paymentsRepository PaymentsRepository,
	usersService UsersService,
	ordersService OrdersService,
	txManager *transaction.Manager,
) *UseCasePayments {
	return &UseCasePayments{
		paymentsRepository: paymentsRepository,
		usersService:       usersService,
		ordersService:      ordersService,
		txManager:          txManager,
	}
}
