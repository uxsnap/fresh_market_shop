package repositoryProducts

import (
	"context"
	"log"

	"github.com/Masterminds/squirrel"
	uuid "github.com/satori/go.uuid"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
	"github.com/uxsnap/fresh_market_shop/backend/internal/repository/postgres/pgEntity"
)

func (r *ProductsRepository) GetOrderedProductsSum(ctx context.Context, orderProducts []entity.OrderProducts) (int64, error) {
	log.Printf("productsRepository.GetOrderedProductsSum")

	productRow := pgEntity.NewProductRow()
	orderProductsRow := pgEntity.NewOrderProductsRow()

	productUids := make([]uuid.UUID, len(orderProducts))
	for i, v := range orderProducts {
		productUids[i] = v.ProductUid
	}

	stmt, args, err := squirrel.Select("sum(p.price * op.count)").
		From(productRow.Table() + " p").
		Join(orderProductsRow.Table() + " op on op.product_uid = p.uid").
		Where(squirrel.Eq{"p.uid": productUids}).
		PlaceholderFormat(squirrel.Dollar).ToSql()

	if err != nil {
		log.Printf("failed to get order products sum: %v", err)
		return -1, nil
	}

	var sum int64

	if err = r.DB().QueryRow(ctx, stmt, args...).Scan(&sum); err != nil {
		log.Printf("failed to scan order products sum: %v", err)
		return -1, nil
	}

	return sum, nil
}
