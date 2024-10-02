package repositoryProducts

import (
	"context"
	"log"

	"github.com/pkg/errors"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
	"github.com/uxsnap/fresh_market_shop/backend/internal/repository/postgres/pgEntity"
)

func (r *ProductsRepository) CreateProduct(ctx context.Context, product entity.Product) error {
	log.Printf("productsRepository.CreateProduct (uid: %s, name: %s)", product.Uid, product.Name)

	productRow := pgEntity.NewProductRow().FromEntity(product)
	if err := r.Create(ctx, productRow); err != nil {
		log.Printf("failed to create product %s: %v", product.Uid, err)
		return errors.WithStack(err)
	}

	return nil
}
