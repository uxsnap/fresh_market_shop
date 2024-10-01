package repositoryProducts

import (
	"context"
	"log"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
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

func (r *ProductsRepository) UpdateProduct(ctx context.Context, product entity.Product) error {
	log.Printf("productsRepository.UpdateProduct (uid: %s, name: %s)", product.Uid, product.Name)

	productRow := pgEntity.NewProductRow().FromEntity(product)
	if err := r.Update(ctx, productRow, productRow.ConditionUidEqual()); err != nil {
		log.Printf("failed to update product %s: %v", product.Uid, err)
		return errors.WithStack(err)
	}

	return nil
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

func (r *ProductsRepository) GetProductsByCategory(ctx context.Context, categoryUid uuid.UUID, limit uint64, offset uint64) ([]entity.Product, error) {
	log.Printf("productsRepository.GetProductsByCategory (category uid: %s)", categoryUid)

	productRow := pgEntity.NewProductRow().FromEntity(entity.Product{CategoryUid: categoryUid})
	rows := pgEntity.NewProductRows()

	if limit == 0 {
		if err := r.GetSome(ctx, productRow, rows, productRow.ConditionCategoryUidEqual()); err != nil {
			log.Printf("failed to get products by category %s: %v", categoryUid, err)
			return nil, errors.WithStack(err)
		}
	} else {
		if err := r.GetWithLimit(ctx, productRow, rows, productRow.ConditionCategoryUidEqual(), limit, offset); err != nil {
			log.Printf("failed to get products by category %s: %v", categoryUid, err)
			return nil, errors.WithStack(err)
		}
	}

	return rows.ToEntity(), nil
}

func (r *ProductsRepository) GetProducts(
	ctx context.Context,
	categoryUid uuid.UUID,
	ccalMin int64,
	ccalMax int64,
	limit uint64,
	offset uint64,
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
