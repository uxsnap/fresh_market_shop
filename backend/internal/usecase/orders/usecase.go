package useCaseOrders

import (
	"github.com/uxsnap/fresh_market_shop/backend/internal/manager/transaction"
)

type UseCaseOrders struct {
	ordersRepository        OrdersRepository
	orderProductsRepository OrderProductsRepository
	productsCountRepository ProductsCountRepository
	productsRepository      ProductsRepository

	paymentsService PaymentsService
	deliveryService DeliveryService

	txManager *transaction.Manager
}

func New(
	ordersRepository OrdersRepository,
	orderProductsRepository OrderProductsRepository,
	productsCountRepository ProductsCountRepository,
	productsRepository ProductsRepository,
	paymentsService PaymentsService,
	deliveryService DeliveryService,

	txManager *transaction.Manager,
) *UseCaseOrders {
	return &UseCaseOrders{
		ordersRepository:        ordersRepository,
		orderProductsRepository: orderProductsRepository,
		productsCountRepository: productsCountRepository,
		productsRepository:      productsRepository,
		paymentsService:         paymentsService,
		deliveryService:         deliveryService,
		txManager:               txManager,
	}
}
