package repositoryProducts

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"github.com/uxsnap/fresh_market_shop/backend/internal/repository/postgres/pgEntity"
)

func (r *ProductsRepository) CreateProductCount(ctx context.Context, productUid uuid.UUID, count int64) error {
	log.Printf("productsRepository.CreateProductCount (product uid: %s, count: %d)", productUid, count)

	if err := r.Create(ctx, pgEntity.NewProductsCountRow(productUid, count)); err != nil {
		log.Printf("failed to create product count: %v", err)
		return errors.WithStack(err)
	}
	return nil
}

func (r *ProductsRepository) UpdateProductCount(ctx context.Context, productUid uuid.UUID, count int64) error {
	log.Printf("productsRepository.UpdateProductCount (product uid: %s, count: %d)", productUid, count)
	productCountRow := pgEntity.NewProductsCountRow(productUid, count)

	if err := r.Update(ctx, productCountRow, productCountRow.ConditionProductUidEqual()); err != nil {
		log.Printf("failed to update product count: %v", err)
		return errors.WithStack(err)
	}
	return nil
}

func (r *ProductsRepository) GetProductCount(ctx context.Context, productUid uuid.UUID) (int64, bool, error) {
	log.Printf("productsRepository.GetProductCount (product uid: %s)", productUid)

	productCountRow := pgEntity.NewProductsCountRow(productUid, 0)

	if err := r.GetOne(ctx, productCountRow, productCountRow.ConditionProductUidEqual()); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return 0, false, nil
		}
		log.Printf("failed to get product count: %v", err)
		return 0, false, errors.WithStack(err)
	}

	return productCountRow.Count(), true, nil
}

func withPrefix(prefix string, fields []string) []string {
	res := make([]string, 0, len(fields))
	for _, f := range fields {
		res = append(res, prefix+"."+f)
	}
	return res
}
