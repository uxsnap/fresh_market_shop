package repositoryProducts

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgtype"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
	"github.com/uxsnap/fresh_market_shop/backend/internal/repository/postgres/pgEntity"
)

func (r *ProductsRepository) GetProductsWithExtra(
	ctx context.Context,
	categoryUid uuid.UUID,
	ccalMin int64,
	ccalMax int64,
	limit uint64,
	offset uint64,
	createdBefore time.Time,
	createdAfter time.Time,
	withCounts bool,
	withPhotos bool,
) ([]entity.ProductWithExtra, error) {
	log.Printf("productsRepository.GetProductsWithExtra (limit: %d, offset: %d )", limit, offset)

	productRow := pgEntity.NewProductRow()
	productsCountRow := pgEntity.NewProductsCountRow(uuid.UUID{}, 0)
	productPhotoRows := pgEntity.NewProductPhotoRows()

	sqlSelectPart := withPrefix("p", productRow.Columns())
	sqlFromPart := fmt.Sprintf("%s as p", productRow.Table())

	if withCounts {
		sqlSelectPart = append(sqlSelectPart, withPrefix("c", productsCountRow.Columns()[1:])...)
		sqlFromPart = sqlFromPart + fmt.Sprintf(" inner join %s as c on p.uid=c.product_uid", productsCountRow.Table())
	}

	if withPhotos {
		sqlSelectPart = append(
			sqlSelectPart,
			" (select jsonb_agg(jsonb_build_object('id', pp.id, 'product_uid', pp.product_uid, 'img_path', pp.img_path)) from product_photos pp where pp.product_uid = p.uid)",
		)
	}

	sql := squirrel.Select(sqlSelectPart...).From(sqlFromPart).PlaceholderFormat(squirrel.Dollar)

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

	res := []entity.ProductWithExtra{}
	jsonPhotosBuf := []byte{}

	valsForScan := productRow.ValuesForScan()
	if withCounts {
		valsForScan = append(valsForScan, productsCountRow.ValuesForScan()[1:]...)
	}
	if withPhotos {
		valsForScan = append(valsForScan, &jsonPhotosBuf)
	}

	for rows.Next() {

		if err := rows.Scan(valsForScan...); err != nil {
			return nil, errors.WithStack(err)
		}

		product := entity.ProductWithExtra{
			Product: productRow.ToEntity(),
		}

		if withCounts {
			product.StockQuantity = productsCountRow.Count()
		}

		if withPhotos {
			if err := productPhotoRows.FromJson(jsonPhotosBuf); err != nil {
				log.Printf("failed to unmarshal photos for product %s: %v", product.Uid, err)
			} else {
				product.Photos = productPhotoRows.ToEntity()
			}
		}

		res = append(res, product)
	}

	return res, nil
}

func (r *ProductsRepository) GetProductsByNameLikeWithExtra(
	ctx context.Context,
	name string,
	limit uint64,
	offset uint64,
	withCounts bool,
	withPhotos bool,
) ([]entity.ProductWithExtra, error) {
	log.Printf("productsRepository.GetProductsByNameLikeWithExtra (name: %s)", name)
	log.Println(withCounts, withPhotos)

	productRow := pgEntity.NewProductRow()
	productsCountRow := pgEntity.NewProductsCountRow(uuid.UUID{}, 0)
	productPhotoRows := pgEntity.NewProductPhotoRows()

	sqlSelectPart := withPrefix("p", productRow.Columns())
	sqlFromPart := fmt.Sprintf("%s as p", productRow.Table())

	if withCounts {
		sqlSelectPart = append(sqlSelectPart, withPrefix("c", productsCountRow.Columns()[1:])...)
		sqlFromPart = sqlFromPart + fmt.Sprintf(" inner join %s as c on p.uid=c.product_uid", productsCountRow.Table())
	}

	if withPhotos {
		sqlSelectPart = append(
			sqlSelectPart,
			" (select jsonb_agg(jsonb_build_object('id', pp.id, 'product_uid', pp.product_uid, 'img_path', pp.img_path)) from product_photos pp where pp.product_uid = p.uid) as photos",
		)
	}

	sql := squirrel.Select(sqlSelectPart...).
		From(sqlFromPart).
		PlaceholderFormat(squirrel.Dollar).
		Where(squirrel.Like{"LOWER(p.name)": "%" + strings.ToLower(name) + "%"})

	if limit != 0 {
		sql = sql.Limit(limit)
	}

	if offset != 0 {
		sql = sql.Offset(offset)
	}

	stmt, args, err := sql.ToSql()

	if err != nil {
		log.Printf("failed to build sql for GetProductsByNameLikeWithExtra name %s: %v", name, err)
		return nil, errors.WithStack(err)
	}

	log.Printf("stmt: %s", stmt)

	rows, err := r.DB().Query(ctx, stmt, args...)
	if err != nil {
		log.Printf("failed to GetProductsByNameLikeWithExtra name %s: %v", name, err)
	}

	res := []entity.ProductWithExtra{}
	jsonPhotosBuf := []byte{}

	valsForScan := productRow.ValuesForScan()
	if withCounts {
		valsForScan = append(valsForScan, productsCountRow.ValuesForScan()[1:]...)
	}
	if withPhotos {
		valsForScan = append(valsForScan, &jsonPhotosBuf)
	}

	for rows.Next() {

		if err := rows.Scan(valsForScan...); err != nil {
			return nil, errors.WithStack(err)
		}

		product := entity.ProductWithExtra{
			Product: productRow.ToEntity(),
		}

		if withCounts {
			product.StockQuantity = productsCountRow.Count()
		}

		if withPhotos {
			if err := productPhotoRows.FromJson(jsonPhotosBuf); err != nil {
				log.Printf("failed to unmarshal photos for product %s: %v", product.Uid, err)
			} else {
				product.Photos = productPhotoRows.ToEntity()
			}
		}

		res = append(res, product)
	}

	return res, nil
}

func (r *ProductsRepository) GetProductsLikeNamesWithLimitOnEachWithExtra(
	ctx context.Context,
	names []string,
	limit uint64,
	offset uint64,
	withCounts bool,
	withPhotos bool,
) ([]entity.ProductWithExtra, error) {
	log.Printf("productsRepository.GetProductsLikeNamesWithLimitOnEachWithExtra (names: %v)", names)

	if len(names) == 0 {
		return nil, nil
	}

	productRow := pgEntity.NewProductRow()
	productsCountRow := pgEntity.NewProductsCountRow(uuid.UUID{}, 0)
	productPhotoRows := pgEntity.NewProductPhotoRows()

	sqlSelectPart := fmt.Sprintf("SELECT %s", strings.Join(withPrefix("p", productRow.Columns()), ","))
	sqlFromPart := fmt.Sprintf("FROM %s p", productRow.Table())

	if withCounts {
		sqlSelectPart += fmt.Sprintf(", %s", strings.Join(withPrefix("c", productsCountRow.Columns()[1:]), ","))
		sqlFromPart += fmt.Sprintf(" inner join %s c on p.uid=c.product_uid", productsCountRow.Table())
	}

	if withPhotos {
		sqlSelectPart += " (select jsonb_agg(jsonb_build_object('id', pp.id, 'product_uid', pp.product_uid, 'img_path', pp.img_path)) from product_photos pp where pp.product_uid = p.uid) as photos"
	}

	stmt := strings.Builder{}
	stmt.WriteString(
		fmt.Sprintf(
			"%s %s WHERE LOWER(p.name) LIKE %s LIMIT %d OFFSET %d\n",
			sqlSelectPart, sqlFromPart, "%$1%", limit, offset,
		),
	)
	for i := 1; i < len(names); i++ {
		stmt.WriteString("UNION\n")
		stmt.WriteString(
			fmt.Sprintf(
				"%s %s WHERE LOWER(p.name) LIKE %s LIMIT %d OFFSET %d\n",
				sqlSelectPart, sqlFromPart, fmt.Sprintf("%%$%d%%", i+1), limit, offset,
			),
		)
	}

	args := make([]interface{}, len(names))
	for i := 0; i < len(names); i++ {
		args[i] = strings.ToLower(names[i])
	}

	rows, err := r.DB().Query(ctx, stmt.String(), args...)
	if err != nil {
		log.Printf("failed to GetProductsLikeNamesWithLimitOnEachWithExtra: %v", err)
		return nil, errors.WithStack(err)
	}

	res := []entity.ProductWithExtra{}
	jsonPhotosBuf := []byte{}

	valsForScan := productRow.ValuesForScan()
	if withCounts {
		valsForScan = append(valsForScan, productsCountRow.ValuesForScan()[1:]...)
	}
	if withPhotos {
		valsForScan = append(valsForScan, &jsonPhotosBuf)
	}

	for rows.Next() {

		if err := rows.Scan(valsForScan...); err != nil {
			return nil, errors.WithStack(err)
		}

		product := entity.ProductWithExtra{
			Product: productRow.ToEntity(),
		}

		if withCounts {
			product.StockQuantity = productsCountRow.Count()
		}

		if withPhotos {
			if err := productPhotoRows.FromJson(jsonPhotosBuf); err != nil {
				log.Printf("failed to unmarshal photos for product %s: %v", product.Uid, err)
			} else {
				product.Photos = productPhotoRows.ToEntity()
			}
		}

		res = append(res, product)
	}

	return res, nil
}
