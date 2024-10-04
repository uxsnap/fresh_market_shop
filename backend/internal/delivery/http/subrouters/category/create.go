package categorySubrouter

import (
	"context"
	"encoding/json"
	"net/http"

	httpEntity "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/entity"
	httpUtils "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/utils"
)

func (h *CategorySubrouter) createCategory(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	var category httpEntity.Category
	if err := json.NewDecoder(r.Body).Decode(&category); err != nil {
		httpUtils.WriteErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	uid, err := h.ProductsService.CreateCategory(ctx, httpEntity.CategoryToEntity(category))
	if err != nil {
		httpUtils.WriteErrorResponse(w, http.StatusInternalServerError, err)
	}

	httpUtils.WriteResponseJson(w, httpEntity.UUID{
		Uid: uid,
	})
}
