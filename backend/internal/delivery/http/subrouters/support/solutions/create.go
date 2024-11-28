package supportSubrouterSolutions

import (
	"context"
	"net/http"

	"github.com/go-chi/chi"
	uuid "github.com/satori/go.uuid"
	httpEntity "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/entity"
	httpUtils "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/utils"
	errorWrapper "github.com/uxsnap/fresh_market_shop/backend/internal/error_wrapper"
)

func (h SupportSubrouterSolutions) CreateTicketSolution(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	ticketUid, err := uuid.FromString(chi.URLParam(r, "ticket_uid"))
	if err != nil {
		httpUtils.WriteErrorResponse(w, http.StatusBadRequest, nil)
		return
	}

	userInfo, err := httpEntity.AuthUserInfoFromContext(r.Context())
	if err != nil {
		httpUtils.WriteErrorResponse(w, http.StatusBadRequest, nil)
		return
	}

	ticket, isFound, err := h.SupportService.GetTicketByUid(ctx, ticketUid)
	if err != nil {
		httpUtils.WriteErrorResponse(w, http.StatusInternalServerError, err)
		return
	}
	if !isFound {
		httpUtils.WriteErrorResponse(w, http.StatusNotFound, nil)
		return
	}

	if !uuid.Equal(userInfo.UserUid, ticket.SolverUid) {
		httpUtils.WriteErrorResponse(
			w, http.StatusForbidden,
			errorWrapper.NewError("permission denied", "нет доступа для создания решения по обращению"),
		)
		return
	}

	var solution httpEntity.SupportTicketSolution
	if err := httpUtils.DecodeJsonRequest(r, &solution); err != nil {
		httpUtils.WriteErrorResponse(w, http.StatusBadRequest, nil)
		return
	}

	solution.TicketUid = ticketUid
	if err := h.SupportService.CreateTicketSolution(ctx, httpEntity.ConvertSupportTicketSolutionToEntity(solution)); err != nil {
		httpUtils.WriteErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
