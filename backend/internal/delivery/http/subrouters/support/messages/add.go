package supportSubrouterMessages

import (
	"context"
	"net/http"

	"github.com/go-chi/chi"
	uuid "github.com/satori/go.uuid"
	httpEntity "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/entity"
	httpUtils "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/utils"
)

func (h SupportSubrouterMessages) AddTicketMessage(w http.ResponseWriter, r *http.Request) {
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

	var message httpEntity.SupportTicketCommentMessage
	if err := httpUtils.DecodeJsonRequest(r, &message); err != nil {
		httpUtils.WriteErrorResponse(w, http.StatusBadRequest, nil)
		return
	}

	message.TicketUid = ticketUid
	message.SenderUid = userInfo.UserUid

	uid, err := h.SupportService.AddTicketMessage(ctx, httpEntity.ConvertSupportTicketCommentMessageToEntity(message))
	if err != nil {
		httpUtils.WriteErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	httpUtils.WriteResponseJson(w, httpEntity.UUID{Uid: uid})
}
