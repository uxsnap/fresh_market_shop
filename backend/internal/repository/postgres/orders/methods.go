package repositoryOrders

import (
	"context"
	"log"

	"github.com/Masterminds/squirrel"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
	errorWrapper "github.com/uxsnap/fresh_market_shop/backend/internal/error_wrapper"
	"github.com/uxsnap/fresh_market_shop/backend/internal/repository/postgres/pgEntity"
)

func (r *OrdersRepository) CreateOrder(ctx context.Context, order entity.Order) error {
	log.Printf("productsRepository.CreateProduct (uid: %s)", order.Uid)

	p, err := pgEntity.NewOrderRow().FromEntity(order)

	if err != nil {
		log.Printf("failed to create order entity %s: %v", order.Uid, err)
		return errorWrapper.NewError(errorWrapper.OrderCreateError, "не удалось создать заказ")
	}

	stmt, args, err := squirrel.
		Insert(p.Table()).
		PlaceholderFormat(squirrel.Dollar).
		Columns("uid", "created_at", "updated_at").
		Values(p.Uid, p.CreatedAt.Time, p.UpdatedAt.Time).ToSql()

	if err != nil {
		log.Printf("failed to create order %s: %v", order.Uid, err)
		return errorWrapper.NewError(errorWrapper.OrderCreateError, "не удалось создать заказ")
	}

	_, err = r.DB().Exec(ctx, stmt, args...)

	if err != nil {
		log.Printf("failed to create order %s: %v", order.Uid, err)
		return errorWrapper.NewError(errorWrapper.OrderCreateError, "не удалось создать заказ")
	}

	return nil
}
