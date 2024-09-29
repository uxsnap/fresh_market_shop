package useCaseCategories

import (
	"context"

	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

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

// переедут в файлы
func (s *UseCaseCategories) GetAllCategories(ctx context.Context) ([]entity.Category, error) {
	return s.categoriesRepository.GetAllCategories(ctx)
}
