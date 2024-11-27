package useCaseSupport

import (
	"context"
	"log"

	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

func (uc *UseCaseSupport) GetTicketsTopicByUid(ctx context.Context, uid uuid.UUID) (entity.SupportTicketsTopic, bool, error) {
	log.Printf("usecaseSupport.GetTicketsTopicByUid: %s", uid)

	topic, isFound, err := uc.repository.GetSupportTicketsTopicByUid(ctx, uid)
	if err != nil {
		log.Printf("failed to get support tickets topic by uid %s: %v", uid, err)
		return entity.SupportTicketsTopic{}, false, errors.WithStack(err)
	}
	return topic, isFound, nil
}
