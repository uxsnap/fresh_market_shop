package useCaseSupport

import (
	"context"
	"log"

	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

func (uc *UseCaseSupport) GetTicketByUid(ctx context.Context, uid uuid.UUID) (entity.SupportTicket, bool, error) {
	log.Printf("usecaseSupport.GetTicketByUid: %s", uid)

	ticket, isFound, err := uc.repository.GetSupportTicketByUid(ctx, uid)
	if err != nil {
		log.Printf("failed to get support ticket by uid %s: %v", uid, err)
		return entity.SupportTicket{}, false, errors.WithStack(err)
	}
	return ticket, isFound, nil
}
