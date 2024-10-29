package searchSubrouter

import (
	"context"
	"net/http"

	httpEntity "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/entity"
	httpUtils "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/utils"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

func (h *SearchSubrouter) searchProducts(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	qFilters, err := entity.NewQueryFiltersParser().WithRequired(
		entity.QueryFieldName,
	).ParseQuery(r.URL.Query())
	if err != nil {
		httpUtils.WriteErrorResponse(w, http.StatusBadRequest, nil)
		return
	}

	if qFilters.WithCounts || qFilters.WithPhotos {
		products, err := h.ProductsService.GetProductsByNameLikeWithExtra(ctx, qFilters.Name, qFilters)
		if err != nil {
			httpUtils.WriteErrorResponse(w, http.StatusInternalServerError, nil)
			return
		}

		resp := make([]httpEntity.ProductWithExtra, 0, len(products))
		for _, product := range products {
			resp = append(resp, httpEntity.ProductWithExtra{
				Product: httpEntity.ProductFromEntity(product.Product),
				Count:   product.StockQuantity,
				Photos:  httpEntity.ProductPhotosFromEntity(product.Photos),
			})
		}

		httpUtils.WriteResponseJson(w, resp)
		return
	}

	products, err := h.ProductsService.GetProductsByNameLike(ctx, qFilters.Name, qFilters)
	if err != nil {
		httpUtils.WriteErrorResponse(w, http.StatusInternalServerError, nil)
		return
	}

	resp := make([]httpEntity.Product, 0, len(products))
	for _, product := range products {
		resp = append(resp, httpEntity.ProductFromEntity(product))
	}

	httpUtils.WriteResponseJson(w, resp)
}
