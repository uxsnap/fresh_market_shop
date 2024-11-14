package paymentsSubrouter

import (
	"context"
	"net/http"

	httpEntity "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/entity"
	httpUtils "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/utils"
	errorWrapper "github.com/uxsnap/fresh_market_shop/backend/internal/error_wrapper"
)

func (h *PaymentsSubrouter) AddUserPaymentCard(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	// TODO: refactor!
	var card httpEntity.UserFullPaymentCard
	if err := httpUtils.DecodeJsonRequest(r, &card); err != nil {
		httpUtils.WriteErrorResponse(w, http.StatusBadRequest, errorWrapper.NewError(errorWrapper.JsonParsingError, err.Error()))
		return
	}

	uid, err := h.PaymentsService.AddUserPaymentCard(ctx, httpEntity.UserFullPaymentCardToEntity(card))
	if err != nil {
		httpUtils.WriteErrorResponse(w, http.StatusInternalServerError, errorWrapper.NewError(errorWrapper.InternalError, err.Error()))
		return
	}

	httpUtils.WriteResponseJson(w, httpEntity.UUID{
		Uid: uid,
	})
}
