package productsSubrouter

import (
	"context"
	"net/http"

	httpEntity "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/entity"
	httpUtils "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/utils"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
	errorWrapper "github.com/uxsnap/fresh_market_shop/backend/internal/error_wrapper"
)

func (h *ProductsSubrouter) GetProducts(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	qFilters, err := entity.NewQueryFiltersParser().
		WithRequired(
			entity.QueryFieldPage,
			entity.QueryFieldLimit,
		).
		WithAllowed(
			entity.QueryFieldName,
			entity.QueryFieldWithCounts,
			entity.QueryFieldWithPhotos,
			entity.QueryFieldCreatedBefore,
			entity.QueryFieldCreatedAfter,
			entity.QueryFieldCategoryUid,
			entity.QueryFieldCcalMin,
			entity.QueryFieldCcalMax,
			entity.QueryFieldCategoryUids,
		).
		ParseQuery(r.URL.Query())
	if err != nil {
		httpUtils.WriteErrorResponse(w, http.StatusBadRequest, errorWrapper.NewError("parsing query params error", err.Error()))
		return
	}

	if qFilters.WithCounts || qFilters.WithPhotos {
		products, err := h.ProductsService.GetProductsWithExtra(ctx, qFilters)
		if err != nil {
			httpUtils.WriteErrorResponse(w, http.StatusInternalServerError, err)
			return
		}

		resp := httpEntity.ProductsWithExtra{}
		for _, product := range products.Products {
			resp.Products = append(resp.Products, httpEntity.ProductWithExtra{
				Product: httpEntity.ProductFromEntity(product.Product),
				Count:   product.StockQuantity,
				Photos:  httpEntity.ProductPhotosFromEntity(product.Photos),
			})
		}

		resp.Total = products.Total

		httpUtils.WriteResponseJson(w, resp)
		return
	}

	products, err := h.ProductsService.GetProducts(ctx, qFilters)
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
