package useCasePayments

import (
	"context"
	"log"

	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

func (uc *UseCasePayments) GetPayment(ctx context.Context, paymentUid uuid.UUID) (entity.Payment, bool, error) {
	log.Printf("usecasePayments.GetPayment: %s", paymentUid)

	payment, isFound, err := uc.paymentsRepository.GetPaymentByUid(ctx, paymentUid)
	if err != nil {
		log.Printf("failed to get payment %s: %v", paymentUid, err)
		return entity.Payment{}, false, errors.WithStack(err)
	}
	return payment, isFound, nil
}
