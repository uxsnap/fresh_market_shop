package repositoryProducts

import (
	"context"
	"fmt"
	"log"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
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

func (r *ProductsRepository) GetProductsWithCounts(
	ctx context.Context,
	categoryUid uuid.UUID,
	ccalMin int64,
	ccalMax int64,
	limit uint64,
	offset uint64,
) ([]entity.ProductWithStockQuantity, error) {
	log.Printf("productsRepository.GetProducts (limit: %d, offset: %d)", limit, offset)

	productRow := pgEntity.NewProductRow()
	productsCountRow := pgEntity.NewProductsCountRow(uuid.UUID{}, 0)

	sql := squirrel.Select(
		append(withPrefix("p", productRow.Columns()), withPrefix("c", productsCountRow.Columns()[1:])...)...).
		From(fmt.Sprintf("%s as p inner joun %s as c on p.uid=c.product_uid", productRow.Table(), productsCountRow.Table())).
		PlaceholderFormat(squirrel.Dollar)

	if !uuid.Equal(categoryUid, uuid.UUID{}) {
		sql = sql.Where(squirrel.Eq{
			"p.category_uid": pgtype.UUID{
				Bytes:  categoryUid,
				Status: pgtype.Present,
			}})
	}

	if ccalMin > 0 {
		sql = sql.Where(squirrel.GtOrEq{
			"p.ccal": ccalMin,
		})
	}
	if ccalMax > 0 {
		sql = sql.Where(
			squirrel.LtOrEq{
				"p.ccal": ccalMax,
			})
	}
	if limit > 0 {
		sql = sql.Limit(limit)
	}

	stmt, args, err := sql.Offset(offset).ToSql()
	if err != nil {
		log.Printf("failed to build sql query: %v", err)
		return nil, errors.WithStack(err)
	}

	rows, err := r.DB().Query(ctx, stmt, args...)
	if err != nil {
		log.Printf("failed to get products: %v", err)
		return nil, errors.WithStack(err)
	}

	res := make([]entity.ProductWithStockQuantity, 0, rows.CommandTag().RowsAffected())
	for rows.Next() {
		if err := rows.Scan(append(productRow.ValuesForScan(), productsCountRow.ValuesForScan()[1:]...)...); err != nil {
			return nil, errors.WithStack(err)
		}

		res = append(res, entity.ProductWithStockQuantity{
			Product:       productRow.ToEntity(),
			StockQuantity: productsCountRow.Count(),
		})
	}

	return res, nil
}

func withPrefix(prefix string, fields []string) []string {
	res := make([]string, 0, len(fields))
	for _, f := range fields {
		res = append(res, prefix+"."+f)
	}
	return res
}
