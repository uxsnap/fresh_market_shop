package useCaseRecipes

type UseCaseRecipes struct {
	recipesRepository RecipesRepository
}

func New(
	recipesRepository RecipesRepository,
) *UseCaseRecipes {
	return &UseCaseRecipes{
		recipesRepository: recipesRepository,
	}
}
