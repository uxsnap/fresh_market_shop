package deliverySubrouter

import (
	"net/http"

	uuid "github.com/satori/go.uuid"
	httpUtils "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/utils"
	errorWrapper "github.com/uxsnap/fresh_market_shop/backend/internal/error_wrapper"
)

func (h *DeliverySubrouter) CalculateDelivery(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userInfo := httpUtils.GetUserInfoFromContext(ctx)

	var req CalculateDeliveryRequest
	if err := httpUtils.DecodeJsonRequest(r, &req); err != nil {
		httpUtils.WriteErrorResponse(w, http.StatusBadRequest, errorWrapper.NewError(errorWrapper.JsonParsingError, err.Error()))
		return
	}

	deliveryPrice, deliveryTime, err := h.DeliveryService.CalculateDelivery(ctx, userInfo.UserUid, req.OrderUid, req.DeliveryAddressUid)
	if err != nil {
		httpUtils.WriteErrorResponse(w, http.StatusInternalServerError, errorWrapper.NewError(errorWrapper.InternalError, err.Error()))
		return
	}

	httpUtils.WriteResponseJson(w, CalculateDeliveryResponse{
		Price: deliveryPrice,
		Time:  int64(deliveryTime),
	})
}

type CalculateDeliveryRequest struct {
	OrderUid           uuid.UUID `json:"orderUid"`
	DeliveryAddressUid uuid.UUID `json:"deliveryAddressUid"`
}

type CalculateDeliveryResponse struct {
	Price int64 `json:"price"`
	Time  int64 `json:"time"`
}
