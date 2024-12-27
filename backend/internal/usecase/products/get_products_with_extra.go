package useCaseProducts

import (
	"context"
	"log"

	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

func (uc *UseCaseProducts) GetProductsWithExtra(ctx context.Context, qFilters entity.QueryFilters) (entity.ProductsWithExtra, error) {
	log.Printf("ucProducts.GetProductsWithExtra")

	categoryUid := qFilters.CategoryUid

	var products entity.ProductsWithExtra

	if !uuid.Equal(categoryUid, uuid.UUID{}) {
		_, categoryFound, err := uc.categoriesRepository.GetCategoryByUid(ctx, categoryUid)
		if err != nil {
			log.Printf("failed to get category %s: %v", categoryUid, err)
		}

		if !categoryFound {
			log.Printf("category %s not found", categoryUid)
			return products, errors.New("category not found")
		}
	}

	if err := uc.txManager.NewPgTransaction().Execute(ctx, func(ctx context.Context) error {
		productsFromRepo, err := uc.productsRepository.GetProductsWithExtra(ctx, qFilters)
		if err != nil {
			log.Printf("failed to get products: %v", err)
			return errors.WithStack(err)
		}

		total, err := uc.productsRepository.GetProductsTotal(ctx)
		if err != nil {
			log.Printf("failed to get products total: %v", err)
			return errors.WithStack(err)
		}

		products.Products = productsFromRepo
		products.Total = total

		return nil
	}); err != nil {
		log.Printf("failed to get products: %v", err)
		return products, errors.WithStack(err)
	}

	return products, nil
}
