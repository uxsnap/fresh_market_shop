package repositoryOrders

import (
	"context"
	"log"

	"github.com/pkg/errors"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
	"github.com/uxsnap/fresh_market_shop/backend/internal/repository/postgres/pgEntity"
)

func (r *OrdersRepository) CreateOrder(ctx context.Context, order entity.Order) error {
	log.Printf("ordersRepository.CreateOrder: uid %s", order.Uid)

	row, err := pgEntity.NewOrderRow().FromEntity(order)
	if err != nil {
		log.Printf("failed to convert recipe: %v", err)
		return errors.WithStack(err)
	}

	if err := r.Create(ctx, row); err != nil {
		log.Printf("failed to create recipe: %v", err)
		return errors.WithStack(err)
	}
	return nil
}
