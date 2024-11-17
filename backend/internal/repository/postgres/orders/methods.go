package repositoryOrders

import (
	"context"
	"log"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
	errorWrapper "github.com/uxsnap/fresh_market_shop/backend/internal/error_wrapper"
	"github.com/uxsnap/fresh_market_shop/backend/internal/repository/postgres/pgEntity"
)

func (r *OrdersRepository) CreateOrder(ctx context.Context, order entity.Order) error {
	log.Printf("productsRepository.CreateProduct (uid: %s)", order.Uid)

	p := pgEntity.NewOrderRow().FromEntity(order)

	stmt, args, err := squirrel.
		Insert(p.Table()).
		PlaceholderFormat(squirrel.Dollar).
		Columns("uid", "user_uid", "created_at", "updated_at").
		Values(p.Uid, p.UserUid, p.CreatedAt.Time, p.UpdatedAt.Time).ToSql()

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

func (r *OrdersRepository) GetOrderByUid(ctx context.Context, uid uuid.UUID) (entity.Order, bool, error) {
	log.Printf("ordersRepository.GetOrderByUid: %s", uid)

	orderRow := pgEntity.NewOrderRow().FromEntity(entity.Order{Uid: uid})
	if err := r.GetOne(ctx, orderRow, orderRow.ConditionUidEqual()); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return entity.Order{}, false, nil
		}
		log.Printf("failed to get order by uid %s: %v", uid, err)
		return entity.Order{}, false, errors.WithStack(err)
	}
	return orderRow.ToEntity(), true, nil
}
