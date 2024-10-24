package repositoryOrders

import (
	"context"
	"log"

	"github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
	"github.com/uxsnap/fresh_market_shop/backend/internal/repository/postgres/pgEntity"
)

func (r *OrdersRepository) CreateOrder(ctx context.Context, order entity.Order) error {
	log.Printf("productsRepository.CreateProduct (uid: %s)", order.Uid)

	p, err := pgEntity.NewOrderRow().FromEntity(order)

	if err != nil {
		log.Printf("failed to create order entity %s: %v", order.Uid, err)
		return errors.WithStack(err)
	}

	stmt, args, err := squirrel.
		Insert(p.Table()).
		PlaceholderFormat(squirrel.Dollar).
		Columns("uid", "sum", "created_at", "updated_at").
		Values(p.Uid, p.Sum, p.CreatedAt.Time, p.UpdatedAt.Time).ToSql()

	if err != nil {
		log.Printf("failed to create order %s: %v", order.Uid, err)
		return err
	}

	_, err = r.DB().Exec(ctx, stmt, args...)

	return err
}
