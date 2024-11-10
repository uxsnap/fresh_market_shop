package deliverySubrouter

import (
	"context"
	"net/http"

	httpEntity "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/entity"
	httpUtils "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/utils"
	errorWrapper "github.com/uxsnap/fresh_market_shop/backend/internal/error_wrapper"
)

func (h *DeliverySubrouter) UpdateDelivery(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	var delivery httpEntity.Delivery
	if err := httpUtils.DecodeJsonRequest(r, &delivery); err != nil {
		httpUtils.WriteErrorResponse(w, http.StatusBadRequest, errorWrapper.NewError(
			errorWrapper.JsonParsingError, "не удалось распарсить тело запроса",
		))
		return
	}

	if err := h.DeliveryService.UpdateDelivery(ctx, httpEntity.DeliveryToEntity(delivery)); err != nil {
		httpUtils.WriteErrorResponse(w, http.StatusInternalServerError, errorWrapper.NewError(
			errorWrapper.InternalError, "не удалось обновить информацию о доставке",
		))
		return
	}

	w.WriteHeader(http.StatusOK)
}
