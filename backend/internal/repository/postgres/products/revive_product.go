package repositoryProducts

import (
	"context"
	"log"

	"github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
	"github.com/uxsnap/fresh_market_shop/backend/internal/repository/postgres/pgEntity"
)

func (r *ProductsRepository) ReviveProduct(ctx context.Context, productUid uuid.UUID) error {
	log.Printf("productsRepository.ReviveProduct uid: %s", productUid)

	productRow := pgEntity.NewProductRow().FromEntity(entity.Product{Uid: productUid})

	stmt, args, err := squirrel.Update(productRow.Table()).
		Set("is_deleted", false).
		PlaceholderFormat(squirrel.Dollar).
		Where(productRow.ConditionUidEqual()).
		ToSql()

	if err != nil {
		log.Printf("failed to revive product %s", productUid)
		return errors.WithStack(err)
	}

	_, err = r.DB().Exec(ctx, stmt, args...)
	if err != nil {
		log.Printf("failed to revive product %s", productUid)
		return errors.WithStack(err)
	}

	return nil
}
