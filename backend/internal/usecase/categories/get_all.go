package useCaseCategories

import (
	"context"
	"log"

	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

// переедут в файлы
func (s *UseCaseCategories) GetAllCategories(ctx context.Context) ([]entity.Category, error) {
	log.Printf("ucCategories.GetAllCategories")
	return s.categoriesRepository.GetAllCategories(ctx)
}
