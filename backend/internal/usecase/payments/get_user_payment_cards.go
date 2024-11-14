package useCasePayments

import (
	"context"
	"log"

	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

func (uc *UseCasePayments) GetUserPaymentCards(ctx context.Context, userUid uuid.UUID) ([]entity.UserPaymentCard, error) {
	log.Printf("usecasePayments.GetUserPaymentCards: %s", userUid)

	cards, err := uc.paymentsRepository.GetUserPaymentCards(ctx, userUid)
	if err != nil {
		log.Printf("failed to get user payment cards by user uid %s: %v", userUid, err)
		return nil, errors.WithStack(err)
	}

	return cards, nil
}

func (uc *UseCasePayments) GetUserFullPaymentCards(ctx context.Context, userUid uuid.UUID) ([]entity.UserFullPaymentCard, error) {
	log.Printf("usecasePayments.GetUserFullPaymentCards: %s", userUid)

	cards, err := uc.paymentsRepository.GetUserFullPaymentCards(ctx, userUid)
	if err != nil {
		log.Printf("failed to get user full payment cards by user uid %s: %v", userUid, err)
		return nil, errors.WithStack(err)
	}

	return cards, nil
}
