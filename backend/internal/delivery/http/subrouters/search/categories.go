package searchSubrouter

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	httpEntity "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/entity"
	httpUtils "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/utils"
)

func (h *SearchSubrouter) searchCategories(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	var (
		err    error
		page   uint64
		limit  uint64
		offset uint64
	)

	reqName := r.URL.Query().Get("name")
	reqPage := r.URL.Query().Get("page")
	reqLimit := r.URL.Query().Get("limit")

	if len(reqName) == 0 {
		httpUtils.WriteErrorResponse(w, http.StatusBadRequest, errors.New("empty name"))
		return
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

	if len(reqLimit) != 0 {
		limit, err = strconv.ParseUint(reqLimit, 10, 64)
		if err != nil {
			httpUtils.WriteErrorResponse(w, http.StatusBadRequest, err)
			return
		}
	}

	offset = (page - 1) * limit

	categories, err := h.ProductsService.GetCategoriesByNameLike(ctx, reqName, limit, offset)
	if err != nil {
		httpUtils.WriteErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	resp := make([]httpEntity.Category, 0, len(categories))
	for _, category := range categories {
		resp = append(resp, httpEntity.CategoryFromEntity(category))
	}

	httpUtils.WriteResponseJson(w, resp)
}
