package useCasePayments

import (
	"context"
	"log"

	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

func (uc *UseCasePayments) GetUserPaymentCardByUid(ctx context.Context, cardUid uuid.UUID) (entity.UserPaymentCard, bool, error) {
	log.Printf("usecasePayments.GetUserPaymentCardByUid: %s", cardUid)

	card, isFound, err := uc.paymentsRepository.GetUserPaymentCardByUid(ctx, cardUid)
	if err != nil {
		log.Printf("failed to get user payment card by uid %s: %v", cardUid, err)
		return entity.UserPaymentCard{}, false, errors.WithStack(err)
	}

	return card, isFound, nil
}

func (uc *UseCasePayments) GetUserFullPaymentCardByUid(ctx context.Context, cardUid uuid.UUID) (entity.UserFullPaymentCard, bool, error) {
	log.Printf("usecasePayments.GetUserFullPaymentCardByUid: %s", cardUid)

	card, isFound, err := uc.paymentsRepository.GetUserFullPaymentCardByUid(ctx, cardUid)
	if err != nil {
		log.Printf("failed to get user full payment card by uid %s: %v", cardUid, err)
		return entity.UserFullPaymentCard{}, false, errors.WithStack(err)
	}

	return card, isFound, nil
}
