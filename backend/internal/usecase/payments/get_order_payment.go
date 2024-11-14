package useCasePayments

import (
	"context"
	"log"

	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

func (uc *UseCasePayments) GetOrderPayment(ctx context.Context, orderUid uuid.UUID) (entity.Payment, bool, error) {
	log.Printf("usecasePayments.GetOrderPayment: %s", orderUid)

	payment, isFound, err := uc.paymentsRepository.GetPaymentByOrderUid(ctx, orderUid)
	if err != nil {
		log.Printf("failed to get payment by order %s: %v", orderUid, err)
		return entity.Payment{}, false, errors.WithStack(err)
	}
	return payment, isFound, nil
}
