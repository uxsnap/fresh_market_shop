package useCaseSupport

import (
	"context"
	"log"
	"time"

	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

func (uc *UseCaseSupport) EditTicket(ctx context.Context, ticket entity.SupportTicket) error {
	log.Printf("usecaseSupport.EditTicket: %s", ticket.Uid)

	if uuid.Equal(ticket.Uid, uuid.UUID{}) {
		return errors.New("пустой uid обращения")
	}
	if !uuid.Equal(ticket.TopicUid, uuid.UUID{}) {
		_, topicIsFound, err := uc.repository.GetSupportTicketsTopicByUid(ctx, ticket.TopicUid)
		if err != nil {
			log.Printf("failed to edit support ticket %s: %v", ticket.Uid, err)
			return errors.WithStack(err)
		}
		if !topicIsFound {
			log.Printf("failed to edit support ticket %s: topic %s not found", ticket.Uid, ticket.TopicUid)
			return errors.New("пустой uid темы обращения")
		}
	}

	if err := uc.txManager.NewPgTransaction().Execute(ctx, func(ctx context.Context) error {
		savedTicket, isFound, err := uc.repository.GetSupportTicketByUid(ctx, ticket.Uid)
		if err != nil {
			return err
		}
		if !isFound {
			return errors.New("обращение не найдено")
		}
		if savedTicket.Status != entity.SupportTicketStatusCreated {
			return errors.Errorf("нельзя редактировать обращение в статусе '%s'", savedTicket.Status)
		}

		if !uuid.Equal(ticket.TopicUid, uuid.UUID{}) && !uuid.Equal(ticket.TopicUid, savedTicket.TopicUid) {
			savedTicket.TopicUid = ticket.TopicUid
		}
		if len(ticket.FromEmail) != 0 && ticket.FromEmail != savedTicket.FromEmail {
			savedTicket.FromEmail = ticket.FromEmail
		}
		if len(ticket.FromPhone) != 0 && ticket.FromPhone != savedTicket.FromPhone {
			savedTicket.FromPhone = ticket.FromPhone
		}
		if len(ticket.Title) != 0 && ticket.Title != savedTicket.Title {
			savedTicket.Title = ticket.Title
		}
		if len(ticket.Description) != 0 && ticket.Description != savedTicket.Description {
			savedTicket.Description = ticket.Description
		}
		savedTicket.UpdatedAt = time.Now().UTC()
		if err := uc.repository.UpdateSupportTicket(ctx, savedTicket); err != nil {
			return err
		}
		return nil
	}); err != nil {
		log.Printf("failed to edit support ticket %s: %v", ticket.Uid, err)
		return errors.WithStack(err)
	}
	return nil
}
