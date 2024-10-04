package categoriesSubrouter

import (
	"context"
	"encoding/json"
	"net/http"

	httpEntity "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/entity"
	httpUtils "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/utils"
)

func (h *CategoriesSubrouter) updateCategory(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	var category httpEntity.Category
	if err := json.NewDecoder(r.Body).Decode(&category); err != nil {
		httpUtils.WriteErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	if err := h.ProductsService.UpdateCategory(ctx, httpEntity.CategoryToEntity(category)); err != nil {
		httpUtils.WriteErrorResponse(w, http.StatusInternalServerError, err)
	}

	w.WriteHeader(http.StatusOK)
}
