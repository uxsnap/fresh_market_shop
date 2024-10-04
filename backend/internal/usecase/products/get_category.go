package useCaseProducts

import (
	"context"
	"log"

	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

func (uc *UseCaseProducts) GetCategoryByUid(ctx context.Context, uid uuid.UUID) (entity.Category, error) {
	log.Printf("ucProducts.GetCategoryByUid: uid %s", uid)

	category, isFound, err := uc.categoriesRepository.GetCategoryByUid(ctx, uid)
	if err != nil {
		log.Printf("failed to get category %s: %v", uid, err)
		return entity.Category{}, errors.WithStack(err)
	}
	if !isFound {
		log.Printf("category %s not found", uid)
		return entity.Category{}, errors.New("category not found")
	}

	return category, nil
}
