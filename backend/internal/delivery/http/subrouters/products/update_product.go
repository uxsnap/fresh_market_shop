package productsSubrouter

import (
	"context"
	"encoding/json"
	"net/http"

	httpEntity "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/entity"
	httpUtils "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/utils"
)

func (h *ProductsSubrouter) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	var product httpEntity.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		httpUtils.WriteErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	if err := h.ProductsService.UpdateProduct(ctx, httpEntity.ProductToEntity(product)); err != nil {
		httpUtils.WriteErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
