package ordersSubrouter

import (
	"context"
	"net/http"

	httpEntity "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/entity"
	httpUtils "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/utils"
	errorWrapper "github.com/uxsnap/fresh_market_shop/backend/internal/error_wrapper"
)

func (h *OrdersSubrouter) CreateOrder(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	var order httpEntity.ProductsCounts
	if err := httpUtils.DecodeJsonRequest(r, &order); err != nil {
		httpUtils.WriteErrorResponse(w, http.StatusBadRequest, errorWrapper.NewError(
			errorWrapper.JsonParsingError, "не удалось распарсить тело запроса",
		))
		return
	}

	uid, err := h.OrdersService.CreateOrder(ctx, httpEntity.ProductsCountsToEntity(order))

	if err != nil {
		httpUtils.WriteErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	httpUtils.WriteResponseJson(w, httpEntity.UUID{
		Uid: uid,
	})
}
