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

func (h *PaymentsSubrouter) GetOrderPayment(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	orderUid, err := uuid.FromString(chi.URLParam(r, "order_uid"))
	if err != nil {
		httpUtils.WriteErrorResponse(w, http.StatusBadRequest, errorWrapper.NewError("bad request: invalid order uid", err.Error()))
		return
	}

	payment, isFound, err := h.PaymentsService.GetOrderPayment(ctx, orderUid)
	if err != nil {
		httpUtils.WriteErrorResponse(w, http.StatusInternalServerError, errorWrapper.NewError(errorWrapper.InternalError, err.Error()))
		return
	}
	if !isFound {
		httpUtils.WriteErrorResponse(w, http.StatusNotFound, errorWrapper.NewError(errorWrapper.InternalError, "payment not found"))
		return
	}

	httpUtils.WriteResponseJson(w, httpEntity.PaymentFromEntity(payment))
}
