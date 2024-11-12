package useCasePayments

import (
	"context"
	"log"

	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

func (uc *UseCasePayments) GetUserPayments(ctx context.Context, userUid uuid.UUID) ([]entity.Payment, error) {
	log.Printf("usecasePayments.GetUserPayments: %s", userUid)

	payments, err := uc.paymentsRepository.GetPaymentsByUserUid(ctx, userUid)
	if err != nil {
		log.Printf("failed to get user %s payments: %v", userUid, err)
		return nil, errors.WithStack(err)
	}
	return payments, nil
}
