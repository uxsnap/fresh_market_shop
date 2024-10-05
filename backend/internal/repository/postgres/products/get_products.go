package repositoryProducts

import (
	"context"
	"log"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
	"github.com/uxsnap/fresh_market_shop/backend/internal/repository/postgres/pgEntity"
)

func (r *ProductsRepository) GetProducts(
	ctx context.Context,
	categoryUid uuid.UUID,
	ccalMin int64,
	ccalMax int64,
	limit uint64,
	offset uint64,
	createdBefore time.Time,
	createdAfter time.Time,
) ([]entity.Product, error) {
	log.Printf("productsRepository.GetProducts (limit: %d, offset: %d)", limit, offset)

	productRow := pgEntity.NewProductRow()
	productRows := pgEntity.NewProductRows()

	sql := squirrel.Select(productRow.Columns()...).From(productRow.Table()).PlaceholderFormat(squirrel.Dollar)

	if !uuid.Equal(categoryUid, uuid.UUID{}) {
		sql = sql.Where(squirrel.Eq{
			"category_uid": pgtype.UUID{
				Bytes:  categoryUid,
				Status: pgtype.Present,
			}})
	}

	if ccalMin > 0 {
		sql = sql.Where(squirrel.GtOrEq{
			"ccal": ccalMin,
		})
	}
	if ccalMax > 0 {
		sql = sql.Where(
			squirrel.LtOrEq{
				"ccal": ccalMax,
			})
	}
	if createdBefore.Unix() != 0 {
		sql = sql.Where(
			squirrel.LtOrEq{
				"created_at": pgtype.Timestamp{
					Time:   createdBefore,
					Status: pgtype.Present,
				},
			})
	}
	if createdAfter.Unix() != 0 {
		sql = sql.Where(
			squirrel.GtOrEq{
				"created_at": pgtype.Timestamp{
					Time:   createdAfter,
					Status: pgtype.Present,
				},
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
