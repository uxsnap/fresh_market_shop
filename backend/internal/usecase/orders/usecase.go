package useCaseOrders

import (
	"github.com/uxsnap/fresh_market_shop/backend/internal/manager/transaction"
)

type UseCaseOrders struct {
	ordersRepository        OrdersRepository
	orderProductsRepository OrderProductsRepository
	productsCountRepository ProductsCountRepository

	txManager *transaction.Manager
}

func New(
	ordersRepository OrdersRepository,
	orderProductsRepository OrderProductsRepository,
	productsCountRepository ProductsCountRepository,
	txManager *transaction.Manager,
) *UseCaseOrders {
	return &UseCaseOrders{
		ordersRepository:        ordersRepository,
		orderProductsRepository: orderProductsRepository,
		productsCountRepository: productsCountRepository,
		txManager:               txManager,
	}
}
