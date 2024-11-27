package useCaseSupport

import (
	"context"
	"log"

	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

func (uc *UseCaseSupport) GetTicketSolution(ctx context.Context, ticketUid uuid.UUID) (entity.SupportTicketSolution, bool, error) {
	log.Printf("usecaseSupport.GetTicketSolution: ticket uid %s", ticketUid)

	solution, isFound, err := uc.repository.GetSupportTicketSolution(ctx, ticketUid)
	if err != nil {
		log.Printf("failed to get support ticket (%s) solution: %v", ticketUid, err)
		return entity.SupportTicketSolution{}, false, errors.WithStack(err)
	}
	return solution, isFound, nil
}
