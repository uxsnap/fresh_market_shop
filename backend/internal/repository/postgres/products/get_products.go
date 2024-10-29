package repositoryProducts

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
	"github.com/uxsnap/fresh_market_shop/backend/internal/repository/postgres/pgEntity"
)

func (r *ProductsRepository) GetProducts(ctx context.Context, qFilters entity.QueryFilters) ([]entity.Product, error) {
	log.Printf("productsRepository.GetProducts (limit: %d, offset: %d)", qFilters.Limit, qFilters.Offset)

	productRow := pgEntity.NewProductRow()
	productRows := pgEntity.NewProductRows()

	sql := squirrel.Select(productRow.Columns()...).From(productRow.Table()).PlaceholderFormat(squirrel.Dollar)

	if !uuid.Equal(qFilters.CategoryUid, uuid.UUID{}) {
		sql = sql.Where(squirrel.Eq{
			"category_uid": pgtype.UUID{
				Bytes:  qFilters.CategoryUid,
				Status: pgtype.Present,
			}})
	}

	if qFilters.CcalMin > 0 {
		sql = sql.Where(squirrel.GtOrEq{
			"ccal": qFilters.CcalMin,
		})
	}
	if qFilters.CcalMax > 0 {
		sql = sql.Where(
			squirrel.LtOrEq{
				"ccal": qFilters.CcalMax,
			})
	}
	if qFilters.CreatedBefore.Unix() > 0 {
		sql = sql.Where(
			squirrel.LtOrEq{
				"created_at": pgtype.Timestamp{
					Time:   qFilters.CreatedBefore,
					Status: pgtype.Present,
				},
			})
	}
	if qFilters.CreatedAfter.Unix() > 0 {
		sql = sql.Where(
			squirrel.GtOrEq{
				"created_at": pgtype.Timestamp{
					Time:   qFilters.CreatedAfter,
					Status: pgtype.Present,
				},
			})
	}

	if qFilters.Limit > 0 {
		sql = sql.Limit(qFilters.Limit)
	}

	stmt, args, err := sql.Offset(qFilters.Offset).ToSql()
	if err != nil {
		log.Printf("failed to build sql query: %v", err)
		return nil, errors.WithStack(err)
	}

	rows, err := r.DB().Query(ctx, stmt, args...)
	if err != nil {
		log.Printf("failed to get products: %v", err)
		return nil, errors.WithStack(err)
	}

	if err := productRows.ScanAll(rows); err != nil {
		log.Printf("failed to scan products: %v", err)
		return nil, errors.WithStack(err)
	}

	return productRows.ToEntity(), nil
}

func (r *ProductsRepository) GetProductByUid(ctx context.Context, uid uuid.UUID) (entity.Product, bool, error) {
	log.Printf("productsRepository.GetProductByUid (uid: %s)", uid)

	productRow := pgEntity.NewProductRow().FromEntity(entity.Product{Uid: uid})
	if err := r.GetOne(ctx, productRow, productRow.ConditionUidEqual()); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return entity.Product{}, false, nil
		}
		log.Printf("failed to get product %s: %v", uid, err)
		return entity.Product{}, false, errors.WithStack(err)
	}

	return productRow.ToEntity(), true, nil
}

func (r *ProductsRepository) GetProductsByNameLike(ctx context.Context, name string, qFilters entity.QueryFilters) ([]entity.Product, error) {
	log.Printf("productsRepository.GetProductsByNameLike (name: %s)", name)

	productRow := pgEntity.NewProductRow().FromEntity(entity.Product{Name: name})
	rows := pgEntity.NewProductRows()

	if qFilters.Limit != 0 {
		if err := r.GetWithLimit(ctx, productRow, rows, productRow.ConditionNameLike(), qFilters.Limit, qFilters.Offset); err != nil {
			log.Printf("failed to get products by name like %s: %v", name, err)
			return nil, errors.WithStack(err)
		}
	} else {
		if err := r.GetSome(ctx, productRow, rows, productRow.ConditionNameLike()); err != nil {
			log.Printf("failed to get products by name like %s: %v", name, err)
			return nil, errors.WithStack(err)
		}
	}

	return rows.ToEntity(), nil
}

func (r *ProductsRepository) GetProductsLikeNamesWithLimitOnEach(ctx context.Context, names []string, qFilters entity.QueryFilters) ([]entity.Product, error) {
	log.Printf("productsRepository.GetProductsLikeNamesWithLimitOnEach (names: %v)", names)

	if len(names) == 0 {
		return nil, nil
	}

	row := pgEntity.NewProductRow()

	stmt := strings.Builder{}
	stmt.WriteString(
		fmt.Sprintf(
			"SELECT %s FROM products WHERE name LIKE %s LIMIT %d OFFSET %d\n",
			strings.Join(row.Columns(), ","), "%$1%", qFilters.Limit, qFilters.Offset,
		),
	)
	for i := 1; i < len(names); i++ {
		stmt.WriteString("UNION\n")
		stmt.WriteString(
			fmt.Sprintf(
				"SELECT %s FROM products WHERE name LIKE %s LIMIT %d OFFSET %d\n",
				strings.Join(row.Columns(), ","), fmt.Sprintf("%%$%d%%", i+1), qFilters.Limit, qFilters.Offset,
			),
		)
	}

	args := make([]interface{}, len(names))
	for i := 0; i < len(names); i++ {
		args[i] = names[i]
	}

	rows, err := r.DB().Query(ctx, stmt.String(), args...)
	if err != nil {
		log.Printf("failed to GetProductsLikeNamesWithLimitOnEach: %v", err)
		return nil, errors.WithStack(err)
	}

	productsRows := pgEntity.NewProductRows()
	if err := productsRows.ScanAll(rows); err != nil {
		log.Printf("failed to GetProductsLikeNamesWithLimitOnEach: %v", err)
		return nil, errors.WithStack(err)
	}

	return productsRows.ToEntity(), nil
}
