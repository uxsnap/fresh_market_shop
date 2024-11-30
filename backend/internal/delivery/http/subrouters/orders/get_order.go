package ordersSubrouter

import (
	"net/http"

	httpEntity "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/entity"
	httpUtils "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/utils"
	errorWrapper "github.com/uxsnap/fresh_market_shop/backend/internal/error_wrapper"
)

func (h *OrdersSubrouter) GetOrder(w http.ResponseWriter, r *http.Request) {
	// ctx := context.Background()

	_, err := httpEntity.AuthUserInfoFromContext(r.Context())
	if err != nil {
		httpUtils.WriteErrorResponse(w, http.StatusBadRequest, errorWrapper.NewError(
			errorWrapper.JwtAuthMiddleware, err.Error(),
		))
		return
	}

	if err != nil {
		httpUtils.WriteErrorResponse(w, http.StatusInternalServerError, errorWrapper.NewError(
			errorWrapper.JsonParsingError, "не удалось получить заказ",
		))
		return
	}

	httpUtils.WriteResponseJson(w, nil)
}
