package repositoryProducts

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgtype"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
	"github.com/uxsnap/fresh_market_shop/backend/internal/repository/postgres/pgEntity"
)

func (r *ProductsRepository) GetProductsWithExtra(ctx context.Context, qFilters entity.QueryFilters) ([]entity.ProductWithExtra, error) {
	log.Printf("productsRepository.GetProductsWithExtra (limit: %d, offset: %d )", qFilters.Limit, qFilters.Offset)

	productRow := pgEntity.NewProductRow()
	productsCountRow := pgEntity.NewProductCountRow(uuid.UUID{}, 0)
	productPhotoRows := pgEntity.NewProductPhotoRows()

	sqlSelectPart := withPrefix("p", productRow.Columns())
	sqlFromPart := fmt.Sprintf("%s as p", productRow.Table())

	if qFilters.WithCounts {
		sqlSelectPart = append(sqlSelectPart, withPrefix("c", productsCountRow.Columns()[1:])...)
		sqlFromPart = sqlFromPart + fmt.Sprintf(" inner join %s as c on p.uid=c.product_uid", productsCountRow.Table())
	}

	if qFilters.WithPhotos {
		sqlSelectPart = append(
			sqlSelectPart,
			" (select jsonb_agg(jsonb_build_object('id', pp.id, 'product_uid', pp.product_uid, 'img_path', pp.img_path)) from product_photos pp where pp.product_uid = p.uid)",
		)
	}

	sql := squirrel.Select(sqlSelectPart...).From(sqlFromPart).PlaceholderFormat(squirrel.Dollar)

	if !uuid.Equal(qFilters.CategoryUid, uuid.UUID{}) {
		sql = sql.Where(squirrel.Eq{
			"p.category_uid": pgtype.UUID{
				Bytes:  qFilters.CategoryUid,
				Status: pgtype.Present,
			}})
	}

	if qFilters.CcalMin > 0 {
		sql = sql.Where(squirrel.GtOrEq{
			"p.ccal": qFilters.CcalMin,
		})
	}
	if qFilters.CcalMax > 0 {
		sql = sql.Where(
			squirrel.LtOrEq{
				"p.ccal": qFilters.CcalMax,
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

	if !uuid.Equal(qFilters.RecipeUid, uuid.UUID{}) {
		recipeRow := pgEntity.NewRecipeRow()

		sql = sql.
			Join(fmt.Sprintf("%v_%v as rs on rs.product_uid = p.uid", recipeRow.Table(), productRow.Table())).
			Join(recipeRow.Table() + " r on r.uid = rs.recipe_uid").
			Where(squirrel.Eq{"r.uid": qFilters.RecipeUid})
	}

	// TODO: add uids to qFilters
	// if len(uuids) > 0 {
	// 	sql = sql.Where(
	// 		squirrel.Eq{
	// 			"p.uid": uuids,
	// 		})
	// }

	if qFilters.Limit > 0 {
		sql = sql.Limit(qFilters.Limit)
	}

	if len(qFilters.CategoryUids) > 0 {
		categoryUidStrs := make([]string, len(qFilters.CategoryUids))

		for ind, v := range qFilters.CategoryUids {
			categoryUidStrs[ind] = fmt.Sprintf("'%v'", v.String())
		}

		sql = sql.OrderBy(
			fmt.Sprintf("CASE WHEN p.category_uid in (%v) THEN 1 ELSE 2 END", strings.Join(categoryUidStrs, ",")),
		)
	}

	if qFilters.WithRandom {
		sql = sql.OrderBy("random()")
	}

	if !uuid.Equal(qFilters.UserUidForOrder, uuid.UUID{}) {
		orderRow := pgEntity.NewOrderRow()
		orderProductRow := pgEntity.NewOrderProductsRow()

		sql = sql.
			Join(fmt.Sprintf(
				// TODO: Убрать статус in_progress, когда будет добавлен 4 этап заказа
				orderRow.Table()+" o on o.user_uid = '%v' and o.status in ('in_progress', 'done')", qFilters.UserUidForOrder,
			)).
			Join(orderProductRow.Table() + " op on op.product_uid = p.uid and o.uid = op.order_uid")
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

	res := []entity.ProductWithExtra{}
	jsonPhotosBuf := []byte{}

	valsForScan := productRow.ValuesForScan()
	if qFilters.WithCounts {
		valsForScan = append(valsForScan, productsCountRow.ValuesForScan()[1:]...)
	}
	if qFilters.WithPhotos {
		valsForScan = append(valsForScan, &jsonPhotosBuf)
	}

	for rows.Next() {

		if err := rows.Scan(valsForScan...); err != nil {
			return nil, errors.WithStack(err)
		}

		product := entity.ProductWithExtra{
			Product: productRow.ToEntity(),
		}

		if qFilters.WithCounts {
			product.StockQuantity = productsCountRow.Count()
		}

		if qFilters.WithPhotos {
			if err := productPhotoRows.FromJson(jsonPhotosBuf); err != nil {
				product.Photos = nil
				// log.Printf("failed to unmarshal photos for product %s: %v", product.Uid, err)
			} else {
				product.Photos = productPhotoRows.ToEntity()
			}
		}

		res = append(res, product)
	}

	return res, nil
}

func (r *ProductsRepository) GetProductsByNameLikeWithExtra(ctx context.Context, name string, qFilters entity.QueryFilters) ([]entity.ProductWithExtra, error) {
	log.Printf("productsRepository.GetProductsByNameLikeWithExtra (name: %s)", name)
	log.Println(qFilters.WithCounts, qFilters.WithPhotos)

	productRow := pgEntity.NewProductRow()
	productsCountRow := pgEntity.NewProductCountRow(uuid.UUID{}, 0)
	productPhotoRows := pgEntity.NewProductPhotoRows()

	sqlSelectPart := withPrefix("p", productRow.Columns())
	sqlFromPart := fmt.Sprintf("%s as p", productRow.Table())

	if qFilters.WithCounts {
		sqlSelectPart = append(sqlSelectPart, withPrefix("c", productsCountRow.Columns()[1:])...)
		sqlFromPart = sqlFromPart + fmt.Sprintf(" inner join %s as c on p.uid=c.product_uid", productsCountRow.Table())
	}

	if qFilters.WithPhotos {
		sqlSelectPart = append(
			sqlSelectPart,
			" (select jsonb_agg(jsonb_build_object('id', pp.id, 'product_uid', pp.product_uid, 'img_path', pp.img_path)) from product_photos pp where pp.product_uid = p.uid) as photos",
		)
	}

	sql := squirrel.Select(sqlSelectPart...).
		From(sqlFromPart).
		PlaceholderFormat(squirrel.Dollar).
		Where(squirrel.Like{"LOWER(p.name)": "%" + strings.ToLower(name) + "%"})

	if qFilters.Limit != 0 {
		sql = sql.Limit(qFilters.Limit)
	}

	if qFilters.Offset != 0 {
		sql = sql.Offset(qFilters.Offset)
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
	if qFilters.WithCounts {
		valsForScan = append(valsForScan, productsCountRow.ValuesForScan()[1:]...)
	}
	if qFilters.WithPhotos {
		valsForScan = append(valsForScan, &jsonPhotosBuf)
	}

	for rows.Next() {

		if err := rows.Scan(valsForScan...); err != nil {
			return nil, errors.WithStack(err)
		}

		product := entity.ProductWithExtra{
			Product: productRow.ToEntity(),
		}

		if qFilters.WithCounts {
			product.StockQuantity = productsCountRow.Count()
		}

		if qFilters.WithPhotos {
			if err := productPhotoRows.FromJson(jsonPhotosBuf); err != nil {
				// log.Printf("failed to unmarshal photos for product %s: %v", product.Uid, err)
			} else {
				product.Photos = productPhotoRows.ToEntity()
			}
		}

		res = append(res, product)
	}

	return res, nil
}

func (r *ProductsRepository) GetProductsLikeNamesWithLimitOnEachWithExtra(ctx context.Context, names []string, qFilters entity.QueryFilters) ([]entity.ProductWithExtra, error) {
	log.Printf("productsRepository.GetProductsLikeNamesWithLimitOnEachWithExtra (names: %v)", names)

	if len(names) == 0 {
		return nil, nil
	}

	productRow := pgEntity.NewProductRow()
	productsCountRow := pgEntity.NewProductCountRow(uuid.UUID{}, 0)
	productPhotoRows := pgEntity.NewProductPhotoRows()

	sqlSelectPart := fmt.Sprintf("SELECT %s", strings.Join(withPrefix("p", productRow.Columns()), ","))
	if len(names) > 1 {
		sqlSelectPart = "(" + sqlSelectPart
	}
	sqlFromPart := fmt.Sprintf("FROM %s p", productRow.Table())

	if qFilters.WithCounts {
		sqlSelectPart += fmt.Sprintf(", %s", strings.Join(withPrefix("c", productsCountRow.Columns()[1:]), ","))
		sqlFromPart += fmt.Sprintf(" inner join %s c on p.uid=c.product_uid", productsCountRow.Table())
	}

	if qFilters.WithPhotos {
		sqlSelectPart += " (select jsonb_agg(jsonb_build_object('id', pp.id, 'product_uid', pp.product_uid, 'img_path', pp.img_path)) from product_photos pp where pp.product_uid = p.uid) as photos"
	}

	stmt := strings.Builder{}
	stmt.WriteString(
		fmt.Sprintf(
			"%s %s WHERE LOWER(p.name) LIKE '%s' LIMIT %d OFFSET %d\n",
			sqlSelectPart, sqlFromPart, fmt.Sprintf("%%%s%%", names[0]), qFilters.LimitOnEach, qFilters.OffsetOnEach,
		),
	)
	for i := 1; i < len(names); i++ {
		stmt.WriteString(")\nUNION\n")
		stmt.WriteString(
			fmt.Sprintf(
				"%s %s WHERE LOWER(p.name) LIKE '%s' LIMIT %d OFFSET %d\n",
				sqlSelectPart, sqlFromPart, fmt.Sprintf("%%%s%%", names[i]), qFilters.LimitOnEach, qFilters.OffsetOnEach,
			),
		)
	}
	if len(names) > 1 {
		stmt.WriteString(")")
	}

	fmt.Println(stmt.String())

	rows, err := r.DB().Query(ctx, stmt.String())
	if err != nil {
		log.Printf("failed to GetProductsLikeNamesWithLimitOnEachWithExtra: %v", err)
		return nil, errors.WithStack(err)
	}

	res := []entity.ProductWithExtra{}
	jsonPhotosBuf := []byte{}

	valsForScan := productRow.ValuesForScan()
	if qFilters.WithCounts {
		valsForScan = append(valsForScan, productsCountRow.ValuesForScan()[1:]...)
	}
	if qFilters.WithPhotos {
		valsForScan = append(valsForScan, &jsonPhotosBuf)
	}

	for rows.Next() {

		if err := rows.Scan(valsForScan...); err != nil {
			return nil, errors.WithStack(err)
		}

		product := entity.ProductWithExtra{
			Product: productRow.ToEntity(),
		}

		if qFilters.WithCounts {
			product.StockQuantity = productsCountRow.Count()
		}

		if qFilters.WithPhotos {
			if err := productPhotoRows.FromJson(jsonPhotosBuf); err != nil {
				// log.Printf("failed to unmarshal photos for product %s: %v", product.Uid, err)
			} else {
				product.Photos = productPhotoRows.ToEntity()
			}
		}

		res = append(res, product)
	}

	return res, nil
}
