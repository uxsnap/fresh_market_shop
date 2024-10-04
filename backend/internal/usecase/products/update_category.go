package useCaseProducts

import (
	"context"
	"log"
	"time"

	"github.com/pkg/errors"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

func (uc *UseCaseProducts) UpdateCategory(ctx context.Context, category entity.Category) error {
	log.Printf("ucProducts.UpdateCategory: uid %s", category.Uid)

	if err := validateCategory(category); err != nil {
		log.Printf("failed to update category %s: %v", category.Uid, err)
		return errors.WithStack(err)
	}

	savedCategory, isFound, err := uc.categoriesRepository.GetCategoryByUid(ctx, category.Uid)
	if err != nil {
		log.Printf("failed to update category %s: %v", category.Uid, err)
		return errors.WithStack(err)
	}

	if !isFound {
		log.Printf("failed to update category %s: %v", category.Uid, err)
		return errors.Errorf("category not found", category.Uid)
	}

	category.CreatedAt = savedCategory.CreatedAt
	category.UpdatedAt = time.Now().UTC()

	if err := uc.categoriesRepository.UpdateCategory(ctx, category); err != nil {
		log.Printf("failed to update category %s: %v", category.Uid, err)
		return errors.WithStack(err)
	}
	return nil
}
