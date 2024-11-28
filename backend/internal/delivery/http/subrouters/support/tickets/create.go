package supportSubrouterTickets

import (
	"context"
	"net/http"

	httpEntity "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/entity"
	httpUtils "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/utils"
)

func (h SupportSubrouterTickets) CreateTicket(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	var userIsAuthorized bool
	userInfo, err := httpEntity.AuthUserInfoFromContext(r.Context())
	if err == nil {
		userIsAuthorized = true
	}

	var ticket httpEntity.SupportTicket
	if err := httpUtils.DecodeJsonRequest(r, &ticket); err != nil {
		httpUtils.WriteErrorResponse(w, http.StatusBadRequest, nil)
		return
	}

	if userIsAuthorized {
		ticket.UserUid = userInfo.UserUid
	}

	uid, err := h.SupportService.CreateTicket(ctx, httpEntity.ConvertSupportTicketToEntity(ticket))
	if err != nil {
		httpUtils.WriteErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	httpUtils.WriteResponseJson(w, httpEntity.UUID{Uid: uid})
}
