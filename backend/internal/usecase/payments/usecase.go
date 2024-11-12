package useCasePayments

import "github.com/uxsnap/fresh_market_shop/backend/internal/manager/transaction"

type UseCasePayments struct {
	paymentsRepository PaymentsRepository
	usersService       UsersService
	txManager          *transaction.Manager
}

func New(
	paymentsRepository PaymentsRepository,
	usersService UsersService,
	txManager *transaction.Manager,
) *UseCasePayments {
	return &UseCasePayments{
		paymentsRepository: paymentsRepository,
		usersService:       usersService,
		txManager:          txManager,
	}
}
