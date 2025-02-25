package deliverySubrouter

import (
	"net/http"

	uuid "github.com/satori/go.uuid"
	httpEntity "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/entity"
	httpUtils "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/utils"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
	errorWrapper "github.com/uxsnap/fresh_market_shop/backend/internal/error_wrapper"
)

func (h *DeliverySubrouter) CalculateDelivery(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userInfo, _ := httpEntity.AuthUserInfoFromContext(ctx)

	var req CalculateDeliveryRequest
	if err := httpUtils.DecodeJsonRequest(r, &req); err != nil {
		httpUtils.WriteErrorResponse(w, http.StatusBadRequest, errorWrapper.NewError(errorWrapper.JsonParsingError, err.Error()))
		return
	}

	var orderPrice int64 = 1

	if req.OrderUid != uuid.Nil {
		order, isFound, err := h.OrdersService.GetOrder(ctx, entity.QueryFilters{
			OrderUid: req.OrderUid,
		})
		if err != nil {
			httpUtils.WriteErrorResponse(w, http.StatusInternalServerError, errorWrapper.NewError(errorWrapper.InternalError, err.Error()))
			return
		}
		if !isFound {
			httpUtils.WriteErrorResponse(w, http.StatusBadRequest, errorWrapper.NewError("bad request", "order not found"))
			return
		}

		orderPrice = order.Sum
	}

	deliveryPrice, deliveryTime, err := h.DeliveryService.CalculateDelivery(ctx, userInfo.UserUid, req.OrderUid, orderPrice, req.DeliveryAddressUid)
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
