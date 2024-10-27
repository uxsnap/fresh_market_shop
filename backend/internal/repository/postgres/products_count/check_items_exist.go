package repositoryProductsCount

import (
	"context"
	"log"

	"github.com/Masterminds/squirrel"
	uuid "github.com/satori/go.uuid"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
	errorWrapper "github.com/uxsnap/fresh_market_shop/backend/internal/error_wrapper"
	"github.com/uxsnap/fresh_market_shop/backend/internal/repository/postgres/pgEntity"
)

func (r *ProductsCountRepository) CheckIfAllItemsExist(ctx context.Context, productsCounts entity.ProductsCounts) *errorWrapper.Error {
	log.Printf("productsRepository.CheckIfAllItemsExist")
	pc := pgEntity.NewProductCountRow(uuid.UUID{}, 0)

	productCountRows := pgEntity.NewProductCountRows()

	productsUidToCountMap := make(map[uuid.UUID]int64)
	uuids := make([]uuid.UUID, len(productsCounts.Products))

	for ind, v := range productsCounts.Products {
		uuids[ind] = v.ProductUid
		productsUidToCountMap[v.ProductUid] = v.Count
	}

	if err := r.GetSome(ctx, pc, productCountRows, squirrel.Eq{"product_uid": uuids}); err != nil {
		log.Printf("failed to get product counts %v", err)
		return errorWrapper.NewError(errorWrapper.ProductCountError, "ошибка получения полей количества продуктов")
	}

	if len(productCountRows.GetRows()) != len(productsCounts.Products) {
		log.Printf("items does not match")
		return errorWrapper.NewError(errorWrapper.OrderCreateValidation, "продукты не существуют, либо их количество не совпадает")
	}

	for _, row := range productCountRows.GetRows() {
		if productsUidToCountMap[row.ProductUid.Bytes] > row.StockQuantity {
			log.Printf("items does not match")
			return errorWrapper.NewError(errorWrapper.OrderCreateValidation, "продукт не существует, либо количество не совпадает")
		}
	}

	return nil
}
