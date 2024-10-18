package searchSubrouter

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	httpEntity "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/entity"
	httpUtils "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/utils"
)

func (h *SearchSubrouter) multipleSearch(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	var (
		err                error
		limitOnProducts    uint64
		offsetOnProducts   uint64
		productsWithCount  bool
		limitOnCategories  uint64
		offsetOnCategories uint64
		page               uint64
	)

	reqName := r.URL.Query().Get("name")
	reqLimitOnProducts := r.URL.Query().Get("limit_on_products")
	reqLimitOnCategories := r.URL.Query().Get("limit_on_categories")
	reqProductsWithCount := r.URL.Query().Get("products_with_count")
	reqPage := r.URL.Query().Get("page")

	if len(reqName) == 0 {
		httpUtils.WriteErrorResponse(w, http.StatusBadRequest, errors.New("empty name"))
		return
	}

	if len(reqLimitOnProducts) != 0 {
		limitOnProducts, err = strconv.ParseUint(reqLimitOnProducts, 10, 64)
		if err != nil {
			httpUtils.WriteErrorResponse(w, http.StatusBadRequest, err)
			return
		}
	}

	if len(reqLimitOnCategories) != 0 {
		limitOnCategories, err = strconv.ParseUint(reqLimitOnCategories, 10, 64)
		if err != nil {
			httpUtils.WriteErrorResponse(w, http.StatusBadRequest, err)
			return
		}
	}

	if len(reqPage) != 0 {
		page, err = strconv.ParseUint(reqPage, 10, 64)
		if err != nil {
			httpUtils.WriteErrorResponse(w, http.StatusBadRequest, err)
			return
		}
	}
	if page == 0 {
		page = 1
	}

	if len(reqProductsWithCount) != 0 {
		productsWithCount, err = strconv.ParseBool(reqProductsWithCount)
		if err != nil {
			httpUtils.WriteErrorResponse(w, http.StatusBadRequest, err)
			return
		}
	}

	offsetOnProducts = (page - 1) * limitOnProducts
	offsetOnCategories = (page - 1) * limitOnCategories

	resp := multipleSearchResponse{}

	categories, err := h.ProductsService.GetCategoriesByNameLike(ctx, reqName, limitOnCategories, offsetOnCategories)
	if err != nil {
		httpUtils.WriteErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	resp.Categories = make([]httpEntity.Category, 0, len(categories))
	for _, category := range categories {
		resp.Categories = append(resp.Categories, httpEntity.CategoryFromEntity(category))
	}

	if productsWithCount {
		products, err := h.ProductsService.GetProductsByNameLikeWithCounts(ctx, reqName, limitOnProducts, offsetOnProducts)
		if err != nil {
			httpUtils.WriteErrorResponse(w, http.StatusInternalServerError, err)
			return
		}

		resp.ProductsWithCounts = make([]httpEntity.ProductWithCount, 0, len(products))
		for _, product := range products {
			resp.ProductsWithCounts = append(resp.ProductsWithCounts, httpEntity.ProductWithCount{
				Product: httpEntity.ProductFromEntity(product.Product),
				Count:   product.StockQuantity,
			})
		}

		httpUtils.WriteResponseJson(w, resp)
		return
	}

	products, err := h.ProductsService.GetProductsByNameLike(ctx, reqName, limitOnProducts, offsetOnProducts)
	if err != nil {
		httpUtils.WriteErrorResponse(w, http.StatusInternalServerError, err)
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
	ProductsWithCounts []httpEntity.ProductWithCount `json:"productsWithCounts,omitempty"`
}
