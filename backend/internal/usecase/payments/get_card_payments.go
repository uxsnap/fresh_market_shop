package useCasePayments

import (
	"context"
	"log"

	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

func (uc *UseCasePayments) GetCardPayments(ctx context.Context, cardUid uuid.UUID) ([]entity.Payment, error) {
	log.Printf("usecasePayments.GetCardPayments: %s", cardUid)

	payments, err := uc.paymentsRepository.GetPaymentsByCardUid(ctx, cardUid)
	if err != nil {
		log.Printf("failed to get card %s payments: %v", cardUid, err)
		return nil, errors.WithStack(err)
	}
	return payments, nil
}
