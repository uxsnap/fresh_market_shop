package useCaseSupport

import (
	"context"
	"log"

	"github.com/pkg/errors"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

func (uc *UseCaseSupport) GetAllTicketsTopics(ctx context.Context) ([]entity.SupportTicketsTopic, error) {
	log.Printf("usecaseSupport.GetAllTicketsTopics")

	topics, err := uc.repository.GetAllSupportTicketsTopics(ctx)
	if err != nil {
		log.Printf("failed to get all support tickets topics: %v", err)
		return nil, errors.WithStack(err)
	}
	return topics, nil
}
