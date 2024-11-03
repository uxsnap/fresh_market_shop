package productsSubrouter

import (
	"context"
	"net/http"

	httpEntity "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/entity"
	httpUtils "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/utils"
	errorWrapper "github.com/uxsnap/fresh_market_shop/backend/internal/error_wrapper"
)

func (h *ProductsSubrouter) CreateProduct(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	var product httpEntity.Product
	if err := httpUtils.DecodeJsonRequest(r, &product); err != nil {
		httpUtils.WriteErrorResponse(w, http.StatusBadRequest, errorWrapper.NewError(
			errorWrapper.JsonParsingError, "не удалось распарсить тело запроса",
		))
		return
	}

	uid, err := h.ProductsService.CreateProduct(ctx, httpEntity.ProductToEntity(product))
	if err != nil {
		httpUtils.WriteErrorResponse(w, http.StatusInternalServerError, errorWrapper.NewError(
			errorWrapper.InternalError, err.Error(),
		))
		return
	}

	httpUtils.WriteResponseJson(w, httpEntity.UUID{
		Uid: uid,
	})
}
