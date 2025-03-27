package ordersSubrouter

import (
	"net/http"

	uuid "github.com/satori/go.uuid"
	httpEntity "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/entity"
	httpUtils "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/utils"
	errorWrapper "github.com/uxsnap/fresh_market_shop/backend/internal/error_wrapper"
)

func (h *OrdersSubrouter) PayOrder(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	userInfo, _ := httpEntity.AuthUserInfoFromContext(ctx)

	var req PayOrderRequest
	if err := httpUtils.DecodeJsonRequest(r, &req); err != nil {
		httpUtils.WriteErrorResponse(w, http.StatusBadRequest, errorWrapper.NewError(
			errorWrapper.JsonParsingError, "не удалось распарсить тело запроса",
		))
		return
	}

	uid, err := h.OrdersService.PayOrder(ctx, userInfo.UserUid, req.OrderUid, req.CardUid, req.DeliveryUid)
	if err != nil {
		httpUtils.WriteErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	httpUtils.WriteResponseJson(w, httpEntity.UUID{Uid: uid})
}

type PayOrderRequest struct {
	OrderUid    uuid.UUID `json:"orderUid"`
	CardUid     uuid.UUID `json:"cardUid"`
	DeliveryUid uuid.UUID `json:"deliveryUid"`
}
