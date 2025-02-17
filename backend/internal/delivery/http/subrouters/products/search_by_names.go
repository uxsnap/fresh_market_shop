package productsSubrouter

import (
	"context"
	"net/http"

	httpEntity "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/entity"
	httpUtils "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/utils"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

const defaultLimitOnEach = 10

func (h *ProductsSubrouter) getProductsByNames(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	var req searchProductsByNamesRequest
	if err := httpUtils.DecodeJsonRequest(r, &req); err != nil {
		httpUtils.WriteErrorResponse(w, http.StatusBadRequest, nil)
		return
	}

	if len(req.Names) == 0 {
		httpUtils.WriteErrorResponse(w, http.StatusBadRequest, nil)
		return
	}

	if req.LimitOnEach == 0 {
		req.LimitOnEach = defaultLimitOnEach
	}
	// TODO: добавить with_photos
	qFilters := entity.QueryFilters{
		LimitOnEach:  req.LimitOnEach,
		OffsetOnEach: req.OffsetOnEach,
		WithCounts:   req.WithCount,
	}

	if req.WithCount {
		products, err := h.ProductsService.GetProductsLikeNamesWithLimitOnEachWithExtra(
			ctx, req.Names, qFilters,
		)
		if err != nil {
			httpUtils.WriteErrorResponse(w, http.StatusInternalServerError, nil)
			return
		}

		resp := make([]httpEntity.ProductWithExtra, 0, len(products))
		for _, product := range products {
			resp = append(resp, httpEntity.ProductWithExtra{
				Product: httpEntity.ProductFromEntity(product.Product),
				Count:   product.StockQuantity,
			})
		}

		httpUtils.WriteResponseJson(w, resp)
		return
	}

	products, err := h.ProductsService.GetProductsLikeNamesWithLimitOnEach(ctx, req.Names, qFilters)
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

type searchProductsByNamesRequest struct {
	Names        []string `json:"names"`
	LimitOnEach  uint64   `json:"limitOnEach"`
	OffsetOnEach uint64   `json:"offsetOnEach"`
	WithCount    bool     `json:"withCount"`
}
