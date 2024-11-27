package useCaseSupport

import (
	"context"
	"log"

	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

func (uc *UseCaseSupport) CreateTicketsTopic(ctx context.Context, topic entity.SupportTicketsTopic) (uuid.UUID, error) {
	log.Printf("usecaseSupport.CreateTicketsTopic: name %s", topic.Name)

	if err := validateTicketsTopic(topic); err != nil {
		log.Printf("failed to create tickets topic: %v", err)
		return uuid.UUID{}, err
	}

	topic.Uid = uuid.NewV4()
	if err := uc.repository.CreateSupportTicketsTopic(ctx, topic); err != nil {
		log.Printf("failed to create support tickets topic:%v", err)
		return uuid.UUID{}, errors.WithStack(err)
	}
	return topic.Uid, nil
}
