package useCaseProducts

import (
	"context"
	"log"

	"github.com/pkg/errors"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

func (uc *UseCaseProducts) GetCategoriesByNameLike(ctx context.Context, name string, limit uint64, offset uint64) ([]entity.Category, error) {
	log.Printf("ucProducts.GetCategoriesByNameLike: name %s", name)

	categories, err := uc.categoriesRepository.GetCategoriesByNameLike(ctx, name, limit, offset)
	if err != nil {
		log.Printf("failed to get categories by name like %s: %v", name, err)
		return nil, errors.WithStack(err)
	}

	return categories, nil
}
