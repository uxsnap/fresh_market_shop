package supportSubrouterTickets

import (
	"context"
	"net/http"

	"github.com/go-chi/chi"
	uuid "github.com/satori/go.uuid"
	httpEntity "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/entity"
	httpUtils "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/utils"
)

func (h SupportSubrouterTickets) TakeTicket(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	userInfo, err := httpEntity.AuthUserInfoFromContext(r.Context())
	if err != nil {
		httpUtils.WriteErrorResponse(w, http.StatusBadRequest, nil)
		return
	}

	// TODO: refactor, add check permissions for support
	if userInfo.Role != "admin" {
		httpUtils.WriteErrorResponse(w, http.StatusForbidden, nil)
		return
	}

	ticketUid, err := uuid.FromString(chi.URLParam(r, "uid"))
	if err != nil {
		httpUtils.WriteErrorResponse(w, http.StatusBadRequest, nil)
		return
	}

	if err := h.SupportService.TakeTicket(ctx, ticketUid, userInfo.UserUid); err != nil {
		httpUtils.WriteErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
