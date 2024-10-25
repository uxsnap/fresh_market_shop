package repositoryProducts

import (
	"context"
	"log"

	sq "github.com/Masterminds/squirrel"
	uuid "github.com/satori/go.uuid"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
	errorWrapper "github.com/uxsnap/fresh_market_shop/backend/internal/error_wrapper"
	"github.com/uxsnap/fresh_market_shop/backend/internal/repository/postgres/pgEntity"
)

func (r *ProductsRepository) CheckIfAllItemsExist(ctx context.Context, ops entity.OrderProducts) *errorWrapper.Error {
	log.Printf("productsRepository.CheckIfAllItemsExist")

	p := pgEntity.NewProductRow()
	pc := pgEntity.NewProductsCountRow(uuid.UUID{}, 0)
	orderProductsRows := pgEntity.NewOrderProductsRows()

	productsUidToCountMap := make(map[uuid.UUID]int64)
	uuids := make([]uuid.UUID, len(ops.Products))

	for ind, v := range ops.Products {
		uuids[ind] = v.Uid
		productsUidToCountMap[v.Uid] = v.Count
	}

	// TODO: Optimize query with counts
	stmt, args, err := sq.Select("p.uid, pc.stock_quantity").
		PlaceholderFormat(sq.Dollar).
		From(p.Table() + " p").
		Join(pc.Table() + " pc on pc.product_uid = p.uid").
		Where(sq.Eq{"uid": uuids}).
		ToSql()

	if err != nil {
		log.Printf("failed to create query for check items %v", err)
		return errorWrapper.NewError(errorWrapper.OrderCreateError, "ошибка создания запроса")
	}

	rows, err := r.DB().Query(ctx, stmt, args...)

	if err != nil {
		log.Printf("failed to query for check items %v", err)
		return errorWrapper.NewError(errorWrapper.OrderCreateError, "ошибка создания запроса")
	}

	orderProductsRows.ScanAll(rows)

	if len(orderProductsRows.Rows) != len(ops.Products) {
		log.Printf("items does not match")
		return errorWrapper.NewError(errorWrapper.OrderCreateValidation, "продукты не существуют, либо их количество не совпадает")
	}

	for _, opRow := range orderProductsRows.Rows {
		if productsUidToCountMap[opRow.Uid.Bytes] > opRow.Count {
			log.Printf("items does not match")
			return errorWrapper.NewError(errorWrapper.OrderCreateValidation, "продукты не существуют, либо их количество не совпадает")
		}
	}

	return nil
}
