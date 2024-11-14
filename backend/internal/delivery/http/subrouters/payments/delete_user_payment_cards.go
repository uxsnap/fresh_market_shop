package paymentsSubrouter

import (
	"context"
	"net/http"

	"github.com/go-chi/chi"
	uuid "github.com/satori/go.uuid"
	httpUtils "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/utils"
	errorWrapper "github.com/uxsnap/fresh_market_shop/backend/internal/error_wrapper"
)

func (h *PaymentsSubrouter) DeleteUserPaymentCards(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	userUid, err := uuid.FromString(chi.URLParam(r, "user_uid"))
	if err != nil {
		httpUtils.WriteErrorResponse(w, http.StatusBadRequest, errorWrapper.NewError("user_uid is required", err.Error()))
		return
	}

	if err := h.PaymentsService.DeleteUserPaymentCards(ctx, userUid); err != nil {
		httpUtils.WriteErrorResponse(w, http.StatusInternalServerError, errorWrapper.NewError(errorWrapper.InternalError, err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
}
