package useCaseRecipes

import "github.com/uxsnap/fresh_market_shop/backend/internal/manager/transaction"

type UseCaseRecipes struct {
	recipesRepository  RecipesRepository
	productsRepository ProductsRepository

	txManager *transaction.Manager
}

func New(
	recipesRepository RecipesRepository,
	productsRepository ProductsRepository,
	txManager *transaction.Manager,
) *UseCaseRecipes {
	return &UseCaseRecipes{
		recipesRepository:  recipesRepository,
		productsRepository: productsRepository,
		txManager:          txManager,
	}
}
