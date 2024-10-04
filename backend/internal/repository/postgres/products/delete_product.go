package repositoryProducts

import (
	"context"
	"log"

	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
	"github.com/uxsnap/fresh_market_shop/backend/internal/repository/postgres/pgEntity"
)

func (r *ProductsRepository) DeleteProduct(ctx context.Context, productUid uuid.UUID) error {
	log.Printf("productsRepository.DeleteProduct uid: %s", productUid)

	productRow := pgEntity.NewProductRow().FromEntity(entity.Product{Uid: productUid})
	if err := r.Delete(ctx, productRow, productRow.ConditionUidEqual()); err != nil {
		log.Printf("failed to delete product %s", productUid)
		return errors.WithStack(err)
	}
	return nil
}
