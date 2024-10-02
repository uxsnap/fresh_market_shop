package useCaseCategories

type UseCaseCategories struct {
	categoriesRepository CategoriesRepository
}

func New(
	categoriesRepository CategoriesRepository,
) *UseCaseCategories {
	return &UseCaseCategories{
		categoriesRepository: categoriesRepository,
	}
}
