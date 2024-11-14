package paymentsSubrouter

import (
	"context"
	"net/http"

	uuid "github.com/satori/go.uuid"
	httpEntity "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/entity"
	httpUtils "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/utils"

	"github.com/go-chi/chi"
	errorWrapper "github.com/uxsnap/fresh_market_shop/backend/internal/error_wrapper"
)

func (h *PaymentsSubrouter) GetUserPaymentCards(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	userUid, err := uuid.FromString(chi.URLParam(r, "user_uid"))
	if err != nil {
		httpUtils.WriteErrorResponse(w, http.StatusBadRequest, errorWrapper.NewError("user_uid is required", err.Error()))
		return
	}

	cards, err := h.PaymentsService.GetUserPaymentCards(ctx, userUid)
	if err != nil {
		httpUtils.WriteErrorResponse(w, http.StatusInternalServerError, errorWrapper.NewError(errorWrapper.InternalError, err.Error()))
		return
	}

	resp := make([]httpEntity.UserPaymentCard, len(cards))
	for i := 0; i < len(cards); i++ {
		resp[i] = httpEntity.UserPaymentCardFromEntity(cards[i])
	}

	httpUtils.WriteResponseJson(w, resp)
}
