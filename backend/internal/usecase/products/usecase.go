package useCaseProducts

import (
	"github.com/uxsnap/fresh_market_shop/backend/internal/manager/transaction"
)

type UseCaseProducts struct {
	productsRepository   ProductsRepository
	categoriesRepository CategoriesRepository

	txManager *transaction.Manager
}

func New(
	productsRepository ProductsRepository,
	categoriesRepository CategoriesRepository,
	txManager *transaction.Manager,
) *UseCaseProducts {
	return &UseCaseProducts{
		productsRepository:   productsRepository,
		categoriesRepository: categoriesRepository,
		txManager:            txManager,
	}
}
