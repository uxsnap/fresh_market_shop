package useCaseProducts

import (
	"context"
	"log"
	"time"

	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

func (uc *UseCaseProducts) CreateCategory(ctx context.Context, category entity.Category) (uuid.UUID, error) {
	log.Printf("ucProducts.CreateCategory: name %s", category.Name)

	if err := validateCategory(category); err != nil {
		log.Printf("failed to create category: %v", err)
		return uuid.UUID{}, errors.WithStack(err)
	}

	category.Uid = uuid.NewV4()
	category.CreatedAt = time.Now().UTC()

	if err := uc.categoriesRepository.CreateCategory(ctx, category); err != nil {
		log.Printf("failed to create category: %v", err)
		return uuid.UUID{}, errors.WithStack(err)
	}
	return category.Uid, nil
}
