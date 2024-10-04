package useCaseProducts

import (
	"context"
	"log"
	"time"

	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

func (s *UseCaseProducts) CreateProduct(ctx context.Context, product entity.Product) (uuid.UUID, error) {
	log.Printf("ucProducts.CreateProduct")

	if err := validateProduct(product); err != nil {
		log.Printf("failed to create product: %v", err)
		return uuid.UUID{}, err
	}

	product.Uid = uuid.NewV4()
	product.CreatedAt = time.Now().UTC()

	if err := s.productsRepository.CreateProduct(ctx, product); err != nil {
		log.Printf("failed to create product: %v", err)
		return uuid.UUID{}, errors.WithStack(err)
	}

	return product.Uid, nil
}
