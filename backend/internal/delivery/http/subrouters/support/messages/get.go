package supportSubrouterMessages

import (
	"context"
	"errors"
	"net/http"

	"github.com/go-chi/chi"
	uuid "github.com/satori/go.uuid"
	httpEntity "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/entity"
	httpUtils "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/utils"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

func (h SupportSubrouterMessages) GetTicketMessages(w http.ResponseWriter, r *http.Request) {
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
		httpUtils.WriteErrorResponse(w, http.StatusNotFound, errors.New("обращение не найдено"))
		return
	}

	if !uuid.Equal(userInfo.UserUid, ticket.UserUid) && !uuid.Equal(userInfo.UserUid, ticket.SolverUid) {
		httpUtils.WriteErrorResponse(w, http.StatusForbidden, errors.New("нет доступа к сообщениям данного обращения"))
		return
	}

	messages, err := h.SupportService.GetTicketMessages(ctx, ticketUid, entity.QueryFilters{})
	if err != nil {
		httpUtils.WriteErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	resp := make([]httpEntity.SupportTicketCommentMessage, len(messages))
	for i := 0; i < len(messages); i++ {
		resp[i] = httpEntity.ConvertSupportTicketCommentMessageFromEntity(messages[i])
	}

	httpUtils.WriteResponseJson(w, resp)
}
