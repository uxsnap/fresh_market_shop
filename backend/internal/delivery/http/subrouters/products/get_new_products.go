package productsSubrouter

import (
	"context"
	"net/http"
	"time"

	httpEntity "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/entity"
	httpUtils "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/utils"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

func (h *ProductsSubrouter) getNewProducts(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	qFilters, err := entity.NewQueryFiltersParser().ParseQuery(r.URL.Query())
	if err != nil {
		httpUtils.WriteErrorResponse(w, http.StatusBadRequest, nil)
		return
	}
	if qFilters.CreatedAfter.Unix() == 0 {
		qFilters.CreatedAfter = time.Now().Add(-time.Hour * 24 * 14).UTC()
	}

	if qFilters.WithCounts || qFilters.WithPhotos {
		products, err := h.ProductsService.GetProductsWithExtra(ctx, qFilters)
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

	products, err := h.ProductsService.GetProducts(ctx, qFilters)
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
