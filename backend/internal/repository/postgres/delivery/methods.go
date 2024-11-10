package repositoryDelivery

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
	"github.com/uxsnap/fresh_market_shop/backend/internal/repository/postgres/pgEntity"
)

func (r *DeliveryRepository) CreateDelivery(ctx context.Context, delivery entity.Delivery) error {
	log.Printf(
		"deliveryRepository.CreateDelivery: delivery uid: %s, order uid: %s",
		delivery.Uid, delivery.OrderUid,
	)

	deliveryRow := pgEntity.NewDeliveryRow().FromEntity(delivery)
	if err := r.Create(ctx, deliveryRow); err != nil {
		log.Printf("failed to create delivery for order %s: %v", delivery.OrderUid, err)
		return errors.WithStack(err)
	}

	return nil
}

func (r *DeliveryRepository) GetDeliveryByOrderUid(ctx context.Context, orderUid uuid.UUID) (entity.Delivery, bool, error) {
	log.Printf("deliveryRepository.GetDeliveryByOrderUid: %s", orderUid)

	deliveryRow := pgEntity.NewDeliveryRow().FromEntity(entity.Delivery{OrderUid: orderUid})

	if err := r.GetOne(ctx, deliveryRow, deliveryRow.ConditionOrderUidEqual()); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return entity.Delivery{}, false, nil
		}

		log.Printf("failed to get delivery by order uid %s: %v", orderUid, err)
		return entity.Delivery{}, false, errors.WithStack(err)
	}

	return deliveryRow.ToEntity(), true, nil
}

func (r *DeliveryRepository) GetDeliveryByUid(ctx context.Context, uid uuid.UUID) (entity.Delivery, bool, error) {
	log.Printf("deliveryRepository.GetDeliveryByUid: %s", uid)

	deliveryRow := pgEntity.NewDeliveryRow().FromEntity(entity.Delivery{Uid: uid})

	if err := r.GetOne(ctx, deliveryRow, deliveryRow.ConditionUidEqual()); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return entity.Delivery{}, false, nil
		}

		log.Printf("failed to get delivery by uid %s: %v", uid, err)
		return entity.Delivery{}, false, errors.WithStack(err)
	}

	return deliveryRow.ToEntity(), true, nil
}

func (r *DeliveryRepository) GetDeliveryHistoryByUser(ctx context.Context, userUid uuid.UUID) ([]entity.Delivery, error) {
	log.Printf("deliveryRepository.GetDeliveryHistoryByUser: uid %s", userUid)

	stmt := fmt.Sprintf(
		`SELECT %s FROM delivery d INNER JOIN orders o ON d.order_uid=o.uid INNER JOIN users u ON o.user_uid=u.uid WHERE o.user_uid=$1`,
		strings.Join(withPrefix("d", pgEntity.NewDeliveryRow().Columns()), ","),
	)

	rows, err := r.DB().Query(ctx, stmt, pgtype.UUID{
		Bytes:  userUid,
		Status: pgtype.Present,
	})
	if err != nil {
		log.Printf("failed to get GetDeliveryHistoryByUser uid %s: %v", userUid, err)
		return nil, errors.WithStack(err)
	}

	row := pgEntity.NewDeliveryRow()

	res := make([]entity.Delivery, 0, 10)
	for rows.Next() {
		if err := row.Scan(rows); err != nil {
			log.Printf("failed to scan delivery: %v", err)
			return nil, errors.WithStack(err)
		}

		res = append(res, row.ToEntity())
	}

	return res, nil
}

func withPrefix(prefix string, fields []string) []string {
	res := make([]string, 0, len(fields))
	for _, f := range fields {
		res = append(res, prefix+"."+f)
	}
	return res
}

func (r *DeliveryRepository) UpdateDelivery(ctx context.Context, delivery entity.Delivery) error {
	log.Printf("deliveryRepository.UpdateDelivery: uid %s", delivery.Uid)

	deliveryRow := pgEntity.NewDeliveryRow().FromEntity(delivery)

	if err := r.Update(ctx, deliveryRow, deliveryRow.ConditionUidEqual()); err != nil {
		log.Printf("failed to update delivery with uid %s: %v", delivery.Uid, err)
		return errors.WithStack(err)
	}
	return nil
}
