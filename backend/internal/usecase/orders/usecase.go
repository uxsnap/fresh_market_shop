package useCaseOrders

import (
	"github.com/uxsnap/fresh_market_shop/backend/internal/manager/transaction"
)

type UseCaseOrders struct {
	ordersRepository   OrdersRepository
	productsRepository ProductsRepository

	txManager *transaction.Manager
}

func New(
	ordersRepository OrdersRepository,
	productsRepository ProductsRepository,
	txManager *transaction.Manager,
) *UseCaseOrders {
	return &UseCaseOrders{
		ordersRepository:   ordersRepository,
		productsRepository: productsRepository,
		txManager:          txManager,
	}
}
