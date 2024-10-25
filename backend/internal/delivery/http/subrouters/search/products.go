package searchSubrouter

import (
	"context"
	"net/http"
	"strconv"

	httpEntity "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/entity"
	httpUtils "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/utils"
)

func (h *SearchSubrouter) searchProducts(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	var (
		withCount  bool
		withPhotos bool
		err        error
		page       int64
		offset     uint64
		limit      uint64
	)

	reqName := r.URL.Query().Get("name")
	reqWithCount := r.URL.Query().Get("with_count")
	reqWithPhotos := r.URL.Query().Get("with_photos")
	reqPage := r.URL.Query().Get("page")
	reqLimit := r.URL.Query().Get("limit")

	if len(reqName) == 0 {
		httpUtils.WriteErrorResponse(w, http.StatusBadRequest, nil)
		return
	}

	if len(reqPage) != 0 {
		page, err = strconv.ParseInt(reqPage, 10, 64)
		if err != nil {
			httpUtils.WriteErrorResponse(w, http.StatusBadRequest, nil)
			return
		}
	}
	if page == 0 {
		page = 1
	}

	if len(reqLimit) != 0 {
		limit, err = strconv.ParseUint(reqLimit, 10, 64)
		if err != nil {
			httpUtils.WriteErrorResponse(w, http.StatusBadRequest, nil)
			return
		}
	}

	offset = uint64((page - 1) * int64(limit))

	if len(reqWithCount) != 0 {
		withCount, err = strconv.ParseBool(reqWithCount)
		if err != nil {
			httpUtils.WriteErrorResponse(w, http.StatusBadRequest, nil)
			return
		}
	}

	if len(reqWithPhotos) != 0 {
		withPhotos, err = strconv.ParseBool(reqWithPhotos)
		if err != nil {
			httpUtils.WriteErrorResponse(w, http.StatusBadRequest, nil)
			return
		}
	}

	if withCount || withPhotos {
		products, err := h.ProductsService.GetProductsByNameLikeWithExtra(ctx, reqName, limit, offset, withCount, withPhotos)
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

	products, err := h.ProductsService.GetProductsByNameLike(ctx, reqName, limit, offset)
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
