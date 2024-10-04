package productsSubrouter

import (
	"context"
	"encoding/json"
	"net/http"

	httpEntity "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/entity"
	httpUtils "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/utils"
)

func (h *ProductsSubrouter) CreateProduct(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	var product httpEntity.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		httpUtils.WriteErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	uid, err := h.ProductsService.CreateProduct(ctx, httpEntity.ProductToEntity(product))
	if err != nil {
		httpUtils.WriteErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	httpUtils.WriteResponseJson(w, httpEntity.UUID{
		Uid: uid,
	})
}
