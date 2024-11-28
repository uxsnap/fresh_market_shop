package supportSubrouterTopics

import (
	"context"
	"net/http"

	httpEntity "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/entity"
	httpUtils "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/utils"
	errorWrapper "github.com/uxsnap/fresh_market_shop/backend/internal/error_wrapper"
)

func (h SupportSubrouterTopics) UpdateTicketsTopic(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	userInfo, err := httpEntity.AuthUserInfoFromContext(r.Context())
	if err != nil {
		httpUtils.WriteErrorResponse(w, http.StatusBadRequest, nil)
		return
	}
	// TODO: refactor
	if userInfo.Role != "admin" {
		httpUtils.WriteErrorResponse(w, http.StatusForbidden, nil)
		return
	}

	var topic httpEntity.SupportTicketsTopic
	if err := httpUtils.DecodeJsonRequest(r, &topic); err != nil {
		httpUtils.WriteErrorResponse(w, http.StatusBadRequest, nil)
		return
	}

	if err := h.SupportService.UpdateTicketsTopic(ctx, httpEntity.ConvertSupportTicketsTopicToEntity(topic)); err != nil {
		httpUtils.WriteErrorResponse(
			w, http.StatusInternalServerError,
			errorWrapper.NewError(errorWrapper.InternalError, err.Error()),
		)
		return
	}

	w.WriteHeader(http.StatusOK)
}
