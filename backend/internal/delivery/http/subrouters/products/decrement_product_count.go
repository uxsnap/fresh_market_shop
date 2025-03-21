package productsSubrouter

import (
	"context"
	"net/http"

	httpEntity "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/entity"
	httpUtils "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/utils"
	errorWrapper "github.com/uxsnap/fresh_market_shop/backend/internal/error_wrapper"
)

func (h *ProductsSubrouter) DecrementProductCount(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	var req httpEntity.ProductCount
	if err := httpUtils.DecodeJsonRequest(r, &req); err != nil {
		httpUtils.WriteErrorResponse(w, http.StatusBadRequest, errorWrapper.NewError(
			errorWrapper.JsonParsingError, "не удалось распарсить тело запроса",
		))
		return
	}

	if err := h.ProductsService.DecrementProductCount(ctx, req.ProductUid, req.Count); err != nil {
		httpUtils.WriteErrorResponse(w, http.StatusInternalServerError, errorWrapper.NewError(
			errorWrapper.JsonParsingError, "не удалось распарсить тело запроса",
		))
		return
	}

	w.WriteHeader(http.StatusOK)
}
