package categorySubrouter

import (
	"context"
	"net/http"

	"github.com/go-chi/chi"
	uuid "github.com/satori/go.uuid"
	httpEntity "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/entity"
	httpUtils "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/utils"
)

func (h *CategorySubrouter) getCategoryByUid(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	categoryUid, err := uuid.FromString(chi.URLParam(r, "uid"))
	if err != nil {
		httpUtils.WriteErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	category, err := h.ProductsService.GetCategoryByUid(ctx, categoryUid)
	if err != nil {
		httpUtils.WriteErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	httpUtils.WriteResponseJson(w, httpEntity.CategoryFromEntity(category))
}
