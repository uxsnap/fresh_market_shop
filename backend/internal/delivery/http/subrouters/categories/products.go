package categoriesSubrouter

import (
	"context"
	"net/http"

	"github.com/go-chi/chi"
	httpEntity "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/entity"
	httpUtils "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/utils"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

func (h *CategoriesSubrouter) getCategoryProducts(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	urlValues := r.URL.Query()
	urlValues.Set(entity.QueryFieldCategoryUid, chi.URLParam(r, "category_uid"))

	qFilters, err := entity.NewQueryFiltersParser().
		WithRequired(entity.QueryFieldCategoryUid).
		WithAllowed(
			entity.QueryFieldWithCounts,
			entity.QueryFieldWithPhotos,
			entity.QueryFieldPage,
			entity.QueryFieldLimit,
			entity.QueryFieldCcalMin,
			entity.QueryFieldCcalMax,
			entity.QueryFieldCreatedBefore,
			entity.QueryFieldCreatedAfter,
		).
		ParseQuery(urlValues)

	if err != nil {
		httpUtils.WriteErrorResponse(w, http.StatusBadRequest, nil)
		return
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
