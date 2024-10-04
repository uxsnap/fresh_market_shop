package useCaseProducts

import (
	"context"
	"log"

	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

func (uc *UseCaseProducts) GetProductByUid(ctx context.Context, productUid uuid.UUID) (entity.Product, bool, error) {
	log.Printf("ucProducts.GetProductByUid: uid: %s", productUid)

	product, isFound, err := uc.productsRepository.GetProductByUid(ctx, productUid)
	return product, isFound, errors.WithStack(err)
}

func (uc *UseCaseProducts) GetProducts(
	ctx context.Context,
	categoryUid uuid.UUID,
	ccalMin int64,
	ccalMax int64,
	limit uint64,
	offset uint64,
) ([]entity.Product, error) {
	log.Printf("ucProducts.GetProducts")

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

	products, err := uc.productsRepository.GetProducts(ctx, categoryUid, ccalMin, ccalMax, limit, offset)
	if err != nil {
		log.Printf("failed to get products: %v", err)
		return nil, errors.WithStack(err)
	}
	return products, nil
}
