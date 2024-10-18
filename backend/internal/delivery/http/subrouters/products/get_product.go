package productsSubrouter

import (
	"context"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	uuid "github.com/satori/go.uuid"
	httpEntity "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/entity"
	httpUtils "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/utils"
)

func (h *ProductsSubrouter) GetProduct(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	var (
		withCount bool
		err error
	)

	uid, err := uuid.FromString(chi.URLParam(r, "uid"))
	if err != nil {
		httpUtils.WriteErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	reqWithCount := r.URL.Query().Get("with_count")
	if len(reqWithCount) != 0 {
		withCount, err = strconv.ParseBool(r.URL.Query().Get("with_count"))
		if err != nil {
			httpUtils.WriteErrorResponse(w, http.StatusBadRequest, err)
			return
		}
	}

	product, isFound, err := h.ProductsService.GetProductByUid(ctx, uid)
	if err != nil {
		httpUtils.WriteErrorResponse(w, http.StatusInternalServerError, err)
		return
	}
	if !isFound {
		httpUtils.WriteErrorResponse(w, http.StatusNotFound, err)
		return
	}

	if !withCount {
		httpUtils.WriteResponseJson(w, httpEntity.ProductFromEntity(product))
		return
	}

	count, _, err := h.ProductsService.GetProductCount(ctx, uid)
	if err != nil {
		httpUtils.WriteErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	httpUtils.WriteResponseJson(w, httpEntity.ProductWithCount{
		Product: httpEntity.ProductFromEntity(product),
		Count:   count,
	})
}
