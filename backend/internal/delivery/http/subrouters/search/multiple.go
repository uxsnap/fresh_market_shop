package searchSubrouter

import (
	"context"
	"fmt"
	"net/http"

	httpEntity "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/entity"
	httpUtils "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/utils"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

func (h *SearchSubrouter) multipleSearch(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	fmt.Println(r.URL.Query())

	qFilters, err := entity.NewQueryFiltersParser().WithRequired(
		entity.QueryFieldName,
	).ParseQuery(r.URL.Query())
	if err != nil {
		httpUtils.WriteErrorResponse(w, http.StatusBadRequest, nil)
		return
	}

	resp := multipleSearchResponse{}

	categories, err := h.ProductsService.GetCategoriesByNameLike(ctx, qFilters.Name, qFilters)
	if err != nil {
		httpUtils.WriteErrorResponse(w, http.StatusInternalServerError, nil)
		return
	}

	resp.Categories = make([]httpEntity.Category, 0, len(categories))
	for _, category := range categories {
		resp.Categories = append(resp.Categories, httpEntity.CategoryFromEntity(category))
	}

	if qFilters.ProductsWithCount || qFilters.ProductsWithPhotos {
		products, err := h.ProductsService.GetProductsByNameLikeWithExtra(ctx, qFilters.Name, qFilters)
		if err != nil {
			httpUtils.WriteErrorResponse(w, http.StatusInternalServerError, nil)
			return
		}

		resp.ProductsWithCounts = make([]httpEntity.ProductWithExtra, 0, len(products))
		for _, product := range products {
			resp.ProductsWithCounts = append(resp.ProductsWithCounts, httpEntity.ProductWithExtra{
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

	resp.Products = make([]httpEntity.Product, 0, len(products))
	for _, product := range products {
		resp.Products = append(resp.Products, httpEntity.ProductFromEntity(product))
	}

	httpUtils.WriteResponseJson(w, resp)
}

type multipleSearchResponse struct {
	Products           []httpEntity.Product          `json:"products,omitempty"`
	Categories         []httpEntity.Category         `json:"categories"`
	ProductsWithCounts []httpEntity.ProductWithExtra `json:"productsWithCounts,omitempty"`
}
