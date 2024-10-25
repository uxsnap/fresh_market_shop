package productsSubrouter

import (
	"context"
	"net/http"
	"strconv"
	"time"

	uuid "github.com/satori/go.uuid"
	httpEntity "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/entity"
	httpUtils "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/utils"
)

func (h *ProductsSubrouter) getNewProducts(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	var (
		err error

		page         int64
		limit        uint64
		ccalMin      int64
		ccalMax      int64
		createdAfter time.Time
		withCounts   bool
		withPhotos   bool
	)

	categoryUid := uuid.FromStringOrNil(r.URL.Query().Get("category_uid"))

	reqPage := r.URL.Query().Get("page")
	if len(reqPage) != 0 {
		page, err = strconv.ParseInt(reqPage, 10, 64)
		if err != nil {
			httpUtils.WriteErrorResponse(w, http.StatusBadRequest, nil)
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
			httpUtils.WriteErrorResponse(w, http.StatusBadRequest, nil)
			return
		}
	}
	if limit == 0 {
		limit = defaultLimit
	}

	reqCcalMin := r.URL.Query().Get("ccalMin")
	if len(reqCcalMin) != 0 {
		ccalMin, err = strconv.ParseInt(reqCcalMin, 10, 64)
		if err != nil {
			httpUtils.WriteErrorResponse(w, http.StatusBadRequest, nil)
			return
		}
	}

	reqCcalMax := r.URL.Query().Get("ccalMax")
	if len(reqCcalMax) != 0 {
		ccalMax, err = strconv.ParseInt(reqCcalMax, 10, 64)
		if err != nil {
			httpUtils.WriteErrorResponse(w, http.StatusBadRequest, nil)
			return
		}
	}

	reqWithCounts := r.URL.Query().Get("with_counts")
	if len(reqWithCounts) != 0 {
		withCounts, err = strconv.ParseBool(reqWithCounts)
		if err != nil {
			httpUtils.WriteErrorResponse(w, http.StatusBadRequest, nil)
			return
		}
	}

	reqWithPhotos := r.URL.Query().Get("with_photos")
	if len(reqWithPhotos) != 0 {
		withPhotos, err = strconv.ParseBool(reqWithPhotos)
		if err != nil {
			httpUtils.WriteErrorResponse(w, http.StatusBadRequest, nil)
			return
		}
	}

	reqCreatedAfter := r.URL.Query().Get("created_after")
	if len(reqCreatedAfter) != 0 {
		createdAfter, err = time.Parse("2006-01-02T15:04:05", reqCreatedAfter)
		if err != nil {
			httpUtils.WriteErrorResponse(w, http.StatusBadRequest, nil)
			return
		}
	} else {
		// по дефолту дата создания <= 2 недель
		createdAfter = time.Now().Add(-time.Hour * 24 * 14).UTC()
	}

	offset := uint64((page - 1) * int64(limit))

	if withCounts || withPhotos {
		products, err := h.ProductsService.GetProductsWithExtra(
			ctx, categoryUid, ccalMin, ccalMax, time.Time{}, createdAfter, limit, offset, withCounts, withPhotos, []uuid.UUID{},
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
				Photos:  httpEntity.ProductPhotosFromEntity(product.Photos),
			})
		}

		httpUtils.WriteResponseJson(w, resp)
		return
	}

	products, err := h.ProductsService.GetProducts(
		ctx, categoryUid, ccalMin, ccalMax, time.Time{}, createdAfter, limit, offset,
	)
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
