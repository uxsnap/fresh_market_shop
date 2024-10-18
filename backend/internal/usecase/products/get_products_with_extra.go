package useCaseProducts

import (
	"context"
	"log"
	"time"

	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

func (uc *UseCaseProducts) GetProductsWithExtra(
	ctx context.Context,
	categoryUid uuid.UUID,
	ccalMin int64,
	ccalMax int64,
	createdBefore time.Time,
	createdAfter time.Time,
	limit uint64,
	offset uint64,
	withCounts bool,
	withPhotos bool,
) ([]entity.ProductWithExtra, error) {
	log.Printf("ucProducts.GetProductsWithExtra")

	if !uuid.Equal(categoryUid, uuid.UUID{}) {
		_, categoryFound, err := uc.categoriesRepository.GetCategoryByUid(ctx, categoryUid)
		if err != nil {
			log.Printf("failed to get category %s: %v", categoryUid, err)
		}

		if !categoryFound {
			log.Printf("category %s not found", categoryUid)
			return nil, errors.New("category not found")
		}
	}

	products, err := uc.productsRepository.GetProductsWithExtra(
		ctx, categoryUid, ccalMin, ccalMax, limit, offset,
		createdBefore, createdAfter, withCounts, withPhotos,
	)
	if err != nil {
		log.Printf("failed to get products: %v", err)
		return nil, errors.WithStack(err)
	}
	return products, nil
}
