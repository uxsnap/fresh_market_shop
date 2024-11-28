package useCaseSupport

import (
	"context"
	"log"
	"time"

	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

func (uc *UseCaseSupport) EditTicketMessage(ctx context.Context, message entity.SupportTicketCommentMessage) error {
	log.Printf("usecaseSupport.EditTicketMessage: ticket uid %s msg uid %s", message.TicketUid, message.Uid)

	if uuid.Equal(message.Uid, uuid.UUID{}) {
		log.Printf("failed to edit support ticket message: empty message uid")
		return errors.New("пустой uid сообщения")
	}
	if len(message.Content) == 0 {
		log.Printf("failed to add support ticket message: empty message content")
		return errors.New("пустое тело сообщения")
	}

	savedMessage, isFound, err := uc.repository.GetSupportTicketCommentMessage(ctx, message.Uid)
	if err != nil {
		log.Printf("failed to edit support ticket message %s: %v", message.Uid, err)
		return errors.WithStack(err)
	}
	if !isFound {
		log.Printf("failed to edit support ticket message %s: message not found", message.Uid)
		return errors.New("сообщение не найдено")
	}

	ticket, isFound, err := uc.repository.GetSupportTicketByUid(ctx, savedMessage.TicketUid)
	if err != nil {
		log.Printf("failed to edit support ticket message %s: %v", message.Uid, err)
		return errors.WithStack(err)
	}
	if !isFound {
		log.Printf("failed to edit support ticket message %s: ticket %s not found", message.Uid, message.TicketUid)
		return errors.New("обращение не найдено")
	}

	if !uuid.Equal(ticket.UserUid, message.SenderUid) && !uuid.Equal(ticket.SolverUid, message.SenderUid) {
		log.Printf("failed to edit support ticket (%s) message: sender is not creator or solver of ticket", message.TicketUid)
		return errors.New("нельзя добавить сообщение по данному обращению")
	}

	switch ticket.Status {
	case entity.SupportTicketStatusSolved, entity.SupportTicketStatusCantSolve:
		log.Printf("failed to add support ticket (%s) message: cant add message for ticket in status '%s'", message.TicketUid, ticket.Status)
		return errors.Errorf("нельзя редактировать сообщения к обращению в статусе '%s'", ticket.Status)
	default:
	}

	savedMessage.Content = message.Content
	savedMessage.UpdatedAt = time.Now().UTC()
	if err := uc.repository.UpdateSupportTicketCommentMessage(ctx, savedMessage); err != nil {
		log.Printf("failed to edit support ticket message %s: %v", message.Uid, err)
		return errors.WithStack(err)
	}
	return nil
}
