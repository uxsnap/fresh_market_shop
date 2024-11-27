package useCaseSupport

import (
	"context"
	"log"

	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

func (uc *UseCaseSupport) GetTicketsSolutionsByTopic(ctx context.Context, topicUid uuid.UUID, qFilters entity.QueryFilters) ([]entity.SupportTicketSolution, error) {
	log.Printf("usecaseSupport.GetTicketsSolutionsByTopic: topic uid %s", topicUid)

	solutions, err := uc.repository.GetSupportTicketSolutionsByTopic(ctx, topicUid, qFilters)
	if err != nil {
		log.Printf("failed to get support tickets solutions by topic (%s): %v", topicUid, err)
		return nil, errors.WithStack(err)
	}
	return solutions, nil
}
