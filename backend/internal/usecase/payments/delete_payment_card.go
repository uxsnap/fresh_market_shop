package useCasePayments

import (
	"context"
	"log"

	uuid "github.com/satori/go.uuid"
)

func (uc *UseCasePayments) DeleteUserPaymentCard(ctx context.Context, cardUid uuid.UUID) error {
	log.Printf("usecasePayments.DeleteUserPaymentCard: %s", cardUid)

	if err := uc.paymentsRepository.DeleteUserPaymentCardByUid(ctx, cardUid); err != nil {
		log.Printf("failed to delete user payment card with uid %s: %v", cardUid, err)
		return err
	}

	return nil
}
