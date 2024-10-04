package useCaseProducts

import (
	"context"
	"log"

	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

func (s *UseCaseProducts) GetAllCategories(ctx context.Context) ([]entity.Category, error) {
	log.Printf("ucProducts.GetAllCategories")
	return s.categoriesRepository.GetAllCategories(ctx)
}
