package categoriesSubrouter

import (
	"context"
	"net/http"

	httpEntity "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/entity"
	httpUtils "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/utils"
)

func (h *CategoriesSubrouter) updateCategory(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	var category httpEntity.Category
	if err := httpUtils.DecodeJsonRequest(r, &category); err != nil {
		httpUtils.WriteErrorResponse(w, http.StatusBadRequest, nil)
		return
	}

	if err := h.ProductsService.UpdateCategory(ctx, httpEntity.CategoryToEntity(category)); err != nil {
		httpUtils.WriteErrorResponse(w, http.StatusInternalServerError, nil)
	}

	w.WriteHeader(http.StatusOK)
}
