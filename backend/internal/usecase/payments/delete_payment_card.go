package useCasePayments

import (
	"context"
	"log"

	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
)

func (uc *UseCasePayments) DeleteUserPaymentCard(ctx context.Context, cardUid uuid.UUID) error {
	log.Printf("usecasePayments.DeleteUserPaymentCard: %s", cardUid)

	if err := uc.txManager.NewPgTransaction().Execute(ctx, func(ctx context.Context) error {
		if err := uc.paymentsRepository.DeleteUserPaymentCardByUid(ctx, cardUid); err != nil {
			log.Printf("failed to delete user payment card with uid %s: %v", cardUid, err)
			return err
		}
		if err := uc.paymentsRepository.DeleteUserFullPaymentCardByUid(ctx, cardUid); err != nil {
			log.Printf("failed to delete user full payment card with uid %s: %v", cardUid, err)
			return err
		}

		return nil
	}); err != nil {
		return errors.WithStack(err)
	}

	return nil
}
