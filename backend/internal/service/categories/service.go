package categoriesService

import (
	"context"

	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

type CategoriesService struct {
	categoriesRepository CategoriesRepository
}

func New(
	categoriesRepository CategoriesRepository,
) *CategoriesService {
	return &CategoriesService{
		categoriesRepository: categoriesRepository,
	}
}

// переедут в файлы
func (s *CategoriesService) GetAllCategories(ctx context.Context) ([]entity.Category, error) {
	return s.categoriesRepository.GetAllCategories(ctx)
}
