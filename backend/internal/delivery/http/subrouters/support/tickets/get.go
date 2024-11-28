package supportSubrouterTickets

import (
	"context"
	"net/http"

	httpEntity "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/entity"
	httpUtils "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/utils"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

func (h SupportSubrouterTickets) GetTickets(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	qFilters, err := entity.NewQueryFiltersParser().WithAllowed(
		entity.QueryFieldUserUid, entity.QueryFieldTopicUid,
		entity.QueryFieldSolverUid, entity.QueryFieldFromEmail,
		entity.QueryFieldStatus, entity.QueryFieldPage, entity.QueryFieldLimit,
	).ParseQuery(r.URL.Query())
	if err != nil {
		httpUtils.WriteErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	tickets, err := h.SupportService.GetTickets(ctx, qFilters)
	if err != nil {
		httpUtils.WriteErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	resp := make([]httpEntity.SupportTicket, len(tickets))
	for i := 0; i < len(tickets); i++ {
		resp[i] = httpEntity.ConvertSupportTicketFromEntity(tickets[i])
	}

	httpUtils.WriteResponseJson(w, resp)
}
