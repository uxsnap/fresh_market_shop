package supportSubrouterTopics

import (
	"context"
	"net/http"

	"github.com/go-chi/chi"
	uuid "github.com/satori/go.uuid"
	httpEntity "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/entity"
	httpUtils "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/utils"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

func (h SupportSubrouterTopics) GetTopicSolutions(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	topicUid, err := uuid.FromString(chi.URLParam(r, "uid"))
	if err != nil {
		httpUtils.WriteErrorResponse(w, http.StatusBadRequest, nil)
		return
	}

	qFilters, err := entity.NewQueryFiltersParser().
		WithAllowed(entity.QueryFieldPage, entity.QueryFieldLimit).
		ParseQuery(r.URL.Query())
	if err != nil {
		httpUtils.WriteErrorResponse(w, http.StatusBadRequest, nil)
		return
	}

	solutions, err := h.SupportService.GetTicketsSolutionsByTopic(ctx, topicUid, qFilters)
	if err != nil {
		httpUtils.WriteErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	resp := make([]httpEntity.SupportTicketSolution, len(solutions))
	for i := 0; i < len(solutions); i++ {
		resp[i] = httpEntity.ConvertSupportTicketSolutionFromEntity(solutions[i])
	}

	httpUtils.WriteResponseJson(w, resp)
}
