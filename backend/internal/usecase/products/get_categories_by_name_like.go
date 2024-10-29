package useCaseProducts

import (
	"context"
	"log"

	"github.com/pkg/errors"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

func (uc *UseCaseProducts) GetCategoriesByNameLike(ctx context.Context, name string, qFilters entity.QueryFilters) ([]entity.Category, error) {
	log.Printf("ucProducts.GetCategoriesByNameLike: name %s", name)

	categories, err := uc.categoriesRepository.GetCategoriesByNameLike(ctx, name, qFilters)
	if err != nil {
		log.Printf("failed to get categories by name like %s: %v", name, err)
		return nil, errors.WithStack(err)
	}

	return categories, nil
}
