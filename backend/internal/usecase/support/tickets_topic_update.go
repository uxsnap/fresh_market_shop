package useCaseSupport

import (
	"context"
	"log"

	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

func (uc *UseCaseSupport) UpdateTicketsTopic(ctx context.Context, topic entity.SupportTicketsTopic) error {
	log.Printf("usecaseSupport.UpdateTicketsTopic: %s", topic.Uid)

	if uuid.Equal(topic.Uid, uuid.UUID{}) {
		log.Printf("failed to update tickets topic: empty uid")
		return errors.New("не указан uid темы обращения")
	}
	if err := validateTicketsTopic(topic); err != nil {
		log.Printf("failed to update tickets topic: %v", err)
		return err
	}

	if err := uc.repository.UpdateSupportTicketsTopic(ctx, topic); err != nil {
		log.Printf("failed to update support tickets topic %s: %v", topic.Uid, err)
		return errors.WithStack(err)
	}
	return nil
}
