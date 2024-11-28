package supportSubrouterTickets

import (
	"context"
	"net/http"

	uuid "github.com/satori/go.uuid"
	httpEntity "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/entity"
	httpUtils "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/utils"
	errorWrapper "github.com/uxsnap/fresh_market_shop/backend/internal/error_wrapper"
)

func (h SupportSubrouterTickets) EditTicket(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	userInfo, err := httpEntity.AuthUserInfoFromContext(r.Context())
	if err != nil {
		httpUtils.WriteErrorResponse(w, http.StatusBadRequest, nil)
		return
	}

	var ticket httpEntity.SupportTicket
	if err := httpUtils.DecodeJsonRequest(r, &ticket); err != nil {
		httpUtils.WriteErrorResponse(w, http.StatusBadRequest, nil)
		return
	}

	if !uuid.Equal(userInfo.UserUid, ticket.UserUid) {
		httpUtils.WriteErrorResponse(
			w, http.StatusForbidden,
			errorWrapper.NewError(errorWrapper.InternalError, "нет доступа редактировать данное обращение"),
		)
		return
	}

	if err := h.SupportService.EditTicket(ctx, httpEntity.ConvertSupportTicketToEntity(ticket)); err != nil {
		httpUtils.WriteErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
