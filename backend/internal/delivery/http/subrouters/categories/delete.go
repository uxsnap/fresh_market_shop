package categoriesSubrouter

import (
	"context"
	"encoding/json"
	"net/http"

	httpEntity "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/entity"
	httpUtils "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/utils"
)

func (h *CategoriesSubrouter) deleteCategory(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	var category httpEntity.UUID
	if err := json.NewDecoder(r.Body).Decode(&category); err != nil {
		httpUtils.WriteErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	if err := h.ProductsService.DeleteCategory(ctx, category.Uid); err != nil {
		httpUtils.WriteErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
