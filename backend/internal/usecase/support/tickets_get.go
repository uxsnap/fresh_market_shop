package useCaseSupport

import (
	"context"
	"log"

	"github.com/pkg/errors"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

func (uc *UseCaseSupport) GetTickets(ctx context.Context, qFilters entity.QueryFilters) ([]entity.SupportTicket, error) {
	log.Printf("usecaseSupport.GetTickets")

	tickets, err := uc.repository.GetSupportTickets(ctx, qFilters)
	if err != nil {
		log.Printf("failed to get support tickets: %v", err)
		return nil, errors.WithStack(err)
	}
	return tickets, nil
}
