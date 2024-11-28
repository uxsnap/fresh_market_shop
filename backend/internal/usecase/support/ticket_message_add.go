package useCaseSupport

import (
	"context"
	"log"
	"time"

	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

func (uc *UseCaseSupport) AddTicketMessage(ctx context.Context, message entity.SupportTicketCommentMessage) (uuid.UUID, error) {
	log.Printf("usecaseSupport.AddTicketMessage: ticket uid %s", message.TicketUid)

	if uuid.Equal(message.TicketUid, uuid.UUID{}) {
		log.Printf("failed to add support ticket message: empty ticket uid")
		return uuid.UUID{}, errors.New("пустой uid обращения")
	}
	if len(message.Content) == 0 {
		log.Printf("failed to add support ticket message: empty message content")
		return uuid.UUID{}, errors.New("пустое тело сообщения")
	}

	ticket, isFound, err := uc.repository.GetSupportTicketByUid(ctx, message.TicketUid)
	if err != nil {
		log.Printf("failed to add support ticket (%s) message: %v", message.TicketUid, err)
		return uuid.UUID{}, errors.WithStack(err)
	}
	if !isFound {
		log.Printf("failed to add support ticket (%s) message: ticket not found", message.TicketUid)
		return uuid.UUID{}, errors.New("обращение не найдено")
	}

	if !uuid.Equal(ticket.UserUid, message.SenderUid) && !uuid.Equal(ticket.SolverUid, message.SenderUid) {
		log.Printf("failed to add support ticket (%s) message: sender is not creator or solver of ticket", message.TicketUid)
		return uuid.UUID{}, errors.New("нельзя добавить сообщение по данному обращению")
	}
	
	switch ticket.Status {
	case entity.SupportTicketStatusSolved, entity.SupportTicketStatusCantSolve:
		log.Printf("failed to add support ticket (%s) message: cant add message for ticket in status '%s'", message.TicketUid, ticket.Status)
		return uuid.UUID{}, errors.Errorf("нельзя добавить сообщение к обращению в статусе '%s'", ticket.Status)
	default:
	}

	message.Uid = uuid.NewV4()
	message.CreatedAt = time.Now().UTC()
	if err := uc.repository.CreateSupportTicketCommentMessage(ctx, message); err != nil {
		log.Printf("failed to add support ticket (%s) message: %v", message.TicketUid, err)
		return uuid.UUID{}, errors.WithStack(err)
	}
	return message.Uid, nil
}
