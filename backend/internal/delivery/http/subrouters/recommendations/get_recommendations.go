package recommendationsSubrouter

import (
	"context"
	"net/http"

	httpEntity "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/entity"
	httpUtils "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/utils"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

func (h *RecommendationsSubrouter) getRecommendations(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	qFilters, err := entity.NewQueryFiltersParser().ParseQuery(r.URL.Query())
	if err != nil {
		httpUtils.WriteErrorResponse(w, http.StatusBadRequest, nil)
		return
	}

	products, err := h.ProductsService.GetProductsWithExtra(
		ctx, qFilters,
	)
	if err != nil {
		httpUtils.WriteErrorResponse(w, http.StatusInternalServerError, nil)
		return
	}

	resp := make([]httpEntity.ProductWithExtra, 0, len(products))
	for _, product := range products {
		resp = append(resp, httpEntity.ProductWithExtra{
			Product: httpEntity.ProductFromEntity(product.Product),
			Photos:  httpEntity.ProductPhotosFromEntity(product.Photos),
		})
	}

	httpUtils.WriteResponseJson(w, resp)
}
