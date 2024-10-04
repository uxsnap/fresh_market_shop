package categoriesSubrouter

import (
	"context"
	"net/http"

	httpEntity "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/entity"
	httpUtils "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/utils"
)

func (h *CategoriesSubrouter) getAllCategories(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	categories, err := h.ProductsService.GetAllCategories(ctx)
	if err != nil {
		httpUtils.WriteErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	response := make([]httpEntity.Category, 0, len(categories))
	for _, category := range categories {
		response = append(response, httpEntity.CategoryFromEntity(category))
	}

	httpUtils.WriteResponseJson(w, response)
}
