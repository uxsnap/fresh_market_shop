package categorySubrouter

import (
	"context"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	uuid "github.com/satori/go.uuid"
	httpEntity "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/entity"
	httpUtils "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/utils"
)

func (h *CategorySubrouter) getCategoryProducts(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	categoryUid, err := uuid.FromString(chi.URLParam(r, "category_uid"))
	if err != nil {
		httpUtils.WriteErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	var (
		page       int64
		limit      uint64
		ccalMin    int64
		ccalMax    int64
		withCounts bool
	)

	reqPage := r.URL.Query().Get("page")
	if len(reqPage) != 0 {
		page, err = strconv.ParseInt(reqPage, 10, 64)
		if err != nil {
			httpUtils.WriteErrorResponse(w, http.StatusBadRequest, err)
			return
		}
	}
	if page <= 0 {
		page = 1
	}

	reqLimit := r.URL.Query().Get("limit")
	if len(reqLimit) != 0 {
		limit, err = strconv.ParseUint(reqLimit, 10, 64)
		if err != nil {
			httpUtils.WriteErrorResponse(w, http.StatusBadRequest, err)
			return
		}
	}

	reqCcalMin := r.URL.Query().Get("ccalMin")
	if len(reqCcalMin) != 0 {
		ccalMin, err = strconv.ParseInt(reqCcalMin, 10, 64)
		if err != nil {
			httpUtils.WriteErrorResponse(w, http.StatusBadRequest, err)
			return
		}
	}

	reqCcalMax := r.URL.Query().Get("ccalMax")
	if len(reqCcalMax) != 0 {
		ccalMax, err = strconv.ParseInt(reqCcalMax, 10, 64)
		if err != nil {
			httpUtils.WriteErrorResponse(w, http.StatusBadRequest, err)
			return
		}
	}

	reqWithCounts := r.URL.Query().Get("with_counts")
	if len(reqWithCounts) != 0 {
		withCounts, err = strconv.ParseBool(reqWithCounts)
		if err != nil {
			httpUtils.WriteErrorResponse(w, http.StatusBadRequest, err)
			return
		}
	}

	offset := uint64((page - 1) * int64(limit))

	if withCounts {
		products, err := h.ProductsService.GetProductsWithCounts(
			ctx, categoryUid, ccalMin, ccalMax, limit, offset,
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

	products, err := h.ProductsService.GetProducts(
		ctx, categoryUid, ccalMin, ccalMax, limit, offset,
	)
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
