package deliverySubrouter

import (
	"context"
	"net/http"

	"github.com/go-chi/chi"
	uuid "github.com/satori/go.uuid"
	httpUtils "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/utils"
	errorWrapper "github.com/uxsnap/fresh_market_shop/backend/internal/error_wrapper"
)

func (h *DeliverySubrouter) GetDeliveryTimeAndPriceForOrder(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	orderUid, err := uuid.FromString(chi.URLParam(r, "order_uid"))
	if err != nil {
		httpUtils.WriteErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	deliveryTime, deliveryPrice, err := h.DeliveryService.GetDeliveryTimeAndPriceForOrder(ctx, orderUid)
	if err != nil {
		httpUtils.WriteErrorResponse(w, http.StatusInternalServerError, errorWrapper.NewError(errorWrapper.InternalError, err.Error()))
		return
	}

	httpUtils.WriteResponseJson(w, DeliveryTimeAndPriceResponse{
		Time:  deliveryTime,
		Price: deliveryPrice,
	})
}

type DeliveryTimeAndPriceResponse struct {
	Time  int64 `json:"time"`
	Price int64 `json:"price"`
}
