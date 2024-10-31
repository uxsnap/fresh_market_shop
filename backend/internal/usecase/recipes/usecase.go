package useCaseRecipes

type UseCaseRecipes struct {
	recipesRepository  RecipesRepository
	productsRepository ProductsRepository
}

func New(
	recipesRepository RecipesRepository,
	productsRepository ProductsRepository,
) *UseCaseRecipes {
	return &UseCaseRecipes{
		recipesRepository:  recipesRepository,
		productsRepository: productsRepository,
	}
}
