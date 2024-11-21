package searchSubrouter

import (
	"context"
	"net/http"

	httpEntity "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/entity"
	httpUtils "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/utils"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

func (h *SearchSubrouter) searchCategories(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	qFilters, err := entity.NewQueryFiltersParser().
		WithRequired(
			entity.QueryFieldName).
		WithAllowed(
			entity.QueryFieldLimit,
			entity.QueryFieldPage).
		ParseQuery(r.URL.Query())
	if err != nil {
		httpUtils.WriteErrorResponse(w, http.StatusBadRequest, nil)
		return
	}

	categories, err := h.ProductsService.GetCategoriesByNameLike(ctx, qFilters.Name, qFilters)
	if err != nil {
		httpUtils.WriteErrorResponse(w, http.StatusInternalServerError, nil)
		return
	}

	resp := make([]httpEntity.Category, 0, len(categories))
	for _, category := range categories {
		resp = append(resp, httpEntity.CategoryFromEntity(category))
	}

	httpUtils.WriteResponseJson(w, resp)
}
