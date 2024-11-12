package repositoryPayments

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
	"github.com/uxsnap/fresh_market_shop/backend/internal/repository/postgres/pgEntity"
)

func (r *PaymentsRepository) CreatePayment(ctx context.Context, payment entity.Payment) error {
	log.Printf("paymentsRepository.CreatePayment: order uid %s", payment.OrderUid)

	paymentRow := pgEntity.NewPaymentRow().FromEntity(payment)
	if err := r.Create(ctx, paymentRow); err != nil {
		log.Printf("failed to create payment for order %s: %v", payment.OrderUid, err)
		return errors.WithStack(err)
	}

	return nil
}

func (r *PaymentsRepository) GetPaymentByUid(ctx context.Context, uid uuid.UUID) (entity.Payment, bool, error) {
	log.Printf("paymentsRepository.GetPaymentByUid: %s", uid)

	paymentRow := pgEntity.NewPaymentRow().FromEntity(entity.Payment{Uid: uid})
	if err := r.GetOne(ctx, paymentRow, paymentRow.ConditionUidEqual()); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return entity.Payment{}, false, nil
		}
		log.Printf("failed to get payment by uid %s: %v", uid, err)
		return entity.Payment{}, false, errors.WithStack(err)
	}
	return paymentRow.ToEntity(), true, nil
}

func (r *PaymentsRepository) GetPaymentByOrderUid(ctx context.Context, orderUid uuid.UUID) (entity.Payment, bool, error) {
	log.Printf("paymentsRepository.GetPaymentByOrderUid: %s", orderUid)

	paymentRow := pgEntity.NewPaymentRow().FromEntity(entity.Payment{OrderUid: orderUid})
	if err := r.GetOne(ctx, paymentRow, paymentRow.ConditionOrderUidEqual()); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return entity.Payment{}, false, nil
		}
		log.Printf("failed to get payment by order uid %s: %v", orderUid, err)
		return entity.Payment{}, false, errors.WithStack(err)
	}
	return paymentRow.ToEntity(), true, nil
}



func (r *PaymentsRepository) GetPaymentsByUserUid(ctx context.Context, userUid uuid.UUID) ([]entity.Payment, error) {
	log.Printf("paymentsRepository.GetPaymentsByUserUid: %s", userUid)

	paymentRow := pgEntity.NewPaymentRow()
	paymentRows := pgEntity.NewPaymentRows()
	if err := r.GetSome(ctx, paymentRow, paymentRows, paymentRow.ConditionUserUidEqual()); err != nil {
		log.Printf("failed to get payments by user uid %s: %v", userUid, err)
		return nil, errors.WithStack(err)
	}

	return paymentRows.ToEntity(), nil
}

