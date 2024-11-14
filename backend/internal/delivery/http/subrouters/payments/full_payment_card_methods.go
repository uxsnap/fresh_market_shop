package paymentsSubrouter

import (
	"context"
	"net/http"

	"github.com/go-chi/chi"
	uuid "github.com/satori/go.uuid"
	httpEntity "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/entity"
	httpUtils "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/utils"
	errorWrapper "github.com/uxsnap/fresh_market_shop/backend/internal/error_wrapper"
)

func (h *PaymentsSubrouter) GetFullPaymentCard(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	cardUid, err := uuid.FromString(chi.URLParam(r, "card_uid"))
	if err != nil {
		httpUtils.WriteErrorResponse(w, http.StatusBadRequest, errorWrapper.NewError("bad request: invalid card uid", err.Error()))
		return
	}

	card, isFound, err := h.PaymentsService.GetUserFullPaymentCardByUid(ctx, cardUid)
	if err != nil {
		httpUtils.WriteErrorResponse(w, http.StatusInternalServerError, errorWrapper.NewError(errorWrapper.InternalError, err.Error()))
		return
	}
	if !isFound {
		httpUtils.WriteErrorResponse(w, http.StatusNotFound, errorWrapper.NewError(errorWrapper.InternalError, "card not found"))
		return
	}

	httpUtils.WriteResponseJson(w, httpEntity.UserFullPaymentCardFromEntity(card))
}

func (h *PaymentsSubrouter) GetUserFullPaymentCards(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	userUid, err := uuid.FromString(chi.URLParam(r, "user_uid"))
	if err != nil {
		httpUtils.WriteErrorResponse(w, http.StatusBadRequest, errorWrapper.NewError("user_uid is required", err.Error()))
		return
	}

	cards, err := h.PaymentsService.GetUserFullPaymentCards(ctx, userUid)
	if err != nil {
		httpUtils.WriteErrorResponse(w, http.StatusInternalServerError, errorWrapper.NewError(errorWrapper.InternalError, err.Error()))
		return
	}

	resp := make([]httpEntity.UserFullPaymentCard, len(cards))
	for i := 0; i < len(cards); i++ {
		resp[i] = httpEntity.UserFullPaymentCardFromEntity(cards[i])
	}

	httpUtils.WriteResponseJson(w, resp)
}
