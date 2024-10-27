package repositoryOrders

import (
	"context"
	"log"

	"github.com/Masterminds/squirrel"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
	errorWrapper "github.com/uxsnap/fresh_market_shop/backend/internal/error_wrapper"
	"github.com/uxsnap/fresh_market_shop/backend/internal/repository/postgres/pgEntity"
)

func (r *OrderProductsRepository) AddOrderProducts(ctx context.Context, orderProducts []entity.OrderProducts) *errorWrapper.Error {
	log.Printf("orderProductsRepository.AddOrderProducts")

	opr := pgEntity.NewOrderProductsRow()

	sql := squirrel.
		Insert(opr.Table()).
		PlaceholderFormat(squirrel.Dollar).
		Columns("order_uid", "product_uid", "count")

	for _, orderProduct := range orderProducts {
		sql = sql.Values(orderProduct.OrderUid, orderProduct.ProductUid, orderProduct.Count)
	}

	stmt, args, err := sql.ToSql()

	if err != nil {
		log.Printf("failed to add order products %s: %v", orderProducts[0].OrderUid, err)
		return errorWrapper.NewError(errorWrapper.OrderCreateError, "не удалось добавить товары заказа")
	}

	_, err = r.DB().Exec(ctx, stmt, args...)

	if err != nil {
		log.Printf("failed to add order products %s: %v", orderProducts[0].OrderUid, err)
		return errorWrapper.NewError(errorWrapper.OrderCreateError, "не удалось добавить товары заказа")
	}

	return nil
}
