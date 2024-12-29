package productsSubrouter

import (
	"context"
	"net/http"

	httpEntity "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/entity"
	httpUtils "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/utils"
	errorWrapper "github.com/uxsnap/fresh_market_shop/backend/internal/error_wrapper"
)

func (h *ProductsSubrouter) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	userInfo, err := httpEntity.AuthUserInfoFromContext(r.Context())
	if err != nil {
		httpUtils.WriteErrorResponse(w, http.StatusBadRequest, errorWrapper.NewError(
			errorWrapper.JsonParsingError, "не удалось найти юзера",
		))
		return
	}

	if userInfo.Role != "admin" {
		httpUtils.WriteErrorResponse(w, http.StatusBadRequest, errorWrapper.NewError(
			errorWrapper.JsonParsingError, "нет разрешений на удаление продукта",
		))
		return
	}

	var product httpEntity.Product
	if err := httpUtils.DecodeJsonRequest(r, &product); err != nil {
		httpUtils.WriteErrorResponse(w, http.StatusBadRequest, nil)
		return
	}

	if err := h.ProductsService.UpdateProduct(ctx, httpEntity.ProductToEntity(product)); err != nil {
		httpUtils.WriteErrorResponse(w, http.StatusInternalServerError, nil)
		return
	}

	w.WriteHeader(http.StatusOK)
}
