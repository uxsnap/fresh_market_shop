package useCaseSupport

import (
	"context"
	"log"

	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

func (uc *UseCaseSupport) GetTicketMessages(ctx context.Context, ticketUid uuid.UUID, qFilters entity.QueryFilters) ([]entity.SupportTicketCommentMessage, error) {
	log.Printf("usecaseSupport.GetTicketMessages: ticket uid %s", ticketUid)

	if uuid.Equal(ticketUid, uuid.UUID{}) {
		log.Printf("failed to get ticket messages: empty ticket uid")
		return nil, errors.New("пустой uid обращения")
	}
	_, ticketIsFound, err := uc.repository.GetSupportTicketByUid(ctx, ticketUid)
	if err != nil {
		log.Printf("failed to get ticket (%s) messages: %v", ticketUid, err)
		return nil, errors.WithStack(err)
	}
	if !ticketIsFound {
		log.Printf("failed to get ticket (%s) messages: ticket not found", ticketUid)
		return nil, errors.New("обращение не найдено")
	}

	messages, err := uc.repository.GetSupportTicketCommentMessages(ctx, ticketUid, qFilters)
	if err != nil {
		log.Printf("failed to get ticket (%s) messages: %v", ticketUid, err)
		return nil, errors.WithStack(err)
	}
	return messages, nil
}
