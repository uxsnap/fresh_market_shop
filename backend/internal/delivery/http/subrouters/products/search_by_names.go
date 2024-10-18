package productsSubrouter

import (
	"context"
	"errors"
	"net/http"

	httpEntity "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/entity"
	httpUtils "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/utils"
)

const defaultLimitOnEach = 10

func (h *ProductsSubrouter) getProductsByNames(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	var req searchProductsByNamesRequest
	if err := httpUtils.DecodeJsonRequest(r, &req); err != nil {
		httpUtils.WriteErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	if len(req.Names) == 0 {
		httpUtils.WriteErrorResponse(w, http.StatusBadRequest, errors.New("empty names"))
		return
	}

	if req.LimitOnEach == 0 {
		req.LimitOnEach = defaultLimitOnEach
	}

	if req.WithCount {
		products, err := h.ProductsService.GetProductsLikeNamesWithLimitOnEachWithCounts(
			ctx, req.Names, req.LimitOnEach, req.OffsetOnEach,
		)
		if err != nil {
			httpUtils.WriteErrorResponse(w, http.StatusInternalServerError, err)
			return
		}

		resp := make([]httpEntity.ProductWithCount, 0, len(products))
		for _, product := range products {
			resp = append(resp, httpEntity.ProductWithCount{
				Product: httpEntity.ProductFromEntity(product.Product),
				Count:   product.StockQuantity,
			})
		}

		httpUtils.WriteResponseJson(w, resp)
		return
	}

	products, err := h.ProductsService.GetProductsLikeNamesWithLimitOnEach(ctx, req.Names, req.LimitOnEach, req.OffsetOnEach)
	if err != nil {
		httpUtils.WriteErrorResponse(w, http.StatusInternalServerError, err)
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
