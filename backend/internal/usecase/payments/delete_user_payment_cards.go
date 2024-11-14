package useCasePayments

import (
	"context"
	"log"

	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
)

func (uc *UseCasePayments) DeleteUserPaymentCards(ctx context.Context, userUid uuid.UUID) error {
	log.Printf("usecasePayments.DeleteAllUserPaymentCards: user uid %s", userUid)

	if err := uc.txManager.NewPgTransaction().Execute(ctx, func(ctx context.Context) error {
		if err := uc.paymentsRepository.DeleteUserPaymentCards(ctx, userUid); err != nil {
			log.Printf("failed to delete user payment cards by user uid %s: %v", userUid, err)
			return errors.WithStack(err)
		}

		if err := uc.paymentsRepository.DeleteUserFullPaymentCards(ctx, userUid); err != nil {
			log.Printf("failed to delete user full payment cards by user uid %s: %v", userUid, err)
			return errors.WithStack(err)
		}

		return nil
	}); err != nil {
		return errors.WithStack(err)
	}

	return nil
}
