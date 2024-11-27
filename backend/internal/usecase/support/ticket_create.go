package useCaseSupport

import (
	"context"
	"log"
	"time"

	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

func (uc *UseCaseSupport) CreateTicket(ctx context.Context, ticket entity.SupportTicket) (uuid.UUID, error) {
	log.Printf("usecaseSupport.CreateTicket: topic_uid %s", ticket.TopicUid)

	if err := validateTicket(ticket); err != nil {
		log.Printf("failed to create support ticket: %v", err)
		return uuid.UUID{}, err
	}

	_, isFound, err := uc.repository.GetSupportTicketsTopicByUid(ctx, ticket.TopicUid)
	if err != nil {
		log.Printf("failed to create support ticket: %v", err)
		return uuid.UUID{}, errors.WithStack(err)
	}
	if !isFound {
		log.Printf("failed to create support ticket: topic not found")
		return uuid.UUID{}, errors.New("тема обращения не найдена")
	}

	ticket.Uid = uuid.NewV4()
	ticket.Status = entity.SupportTicketStatusCreated
	ticket.CreatedAt = time.Now().UTC()

	if err := uc.repository.CreateSupportTicket(ctx, ticket); err != nil {
		log.Printf("failed to create support ticket: %v", err)
		return uuid.UUID{}, errors.WithStack(err)
	}
	return ticket.Uid, nil
}
