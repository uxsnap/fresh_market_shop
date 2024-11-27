package useCaseSupport

import (
	"context"
	"log"

	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
)

func (uc *UseCaseSupport) DeleteTicketsTopic(ctx context.Context, uid uuid.UUID) error {
	log.Printf("usecaseSupport.DeleteTicketsTopic: %s", uid)

	if err := uc.repository.DeleteSupportTicketsTopic(ctx, uid); err != nil {
		log.Printf("failed to delete support tickets topic %s: %v", uid, err)
		return errors.WithStack(err)
	}
	return nil
}
