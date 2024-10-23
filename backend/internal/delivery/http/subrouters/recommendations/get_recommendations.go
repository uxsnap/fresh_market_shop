package recommendationsSubrouter

import (
	"context"
	"net/http"
	"strconv"
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/uxsnap/fresh_market_shop/backend/internal/consts"
	httpEntity "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/entity"
	httpUtils "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/utils"
)

func (h *RecommendationsSubrouter) getRecommendations(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	rQuery := r.URL.Query()

	var (
		limit      = consts.DEFAULT_LIMIT
		withPhotos bool
		err        error
	)

	// userUid := uuid.FromStringOrNil(r.URL.Query().Get("user_uid"))
	categoryUid := uuid.FromStringOrNil(rQuery.Get("category_uid"))

	reqLimit := rQuery.Get("limit")
	if reqLimit != "" {
		limit, err = strconv.ParseUint(reqLimit, 10, 64)
		if err != nil {
			httpUtils.WriteErrorResponse(w, http.StatusBadRequest, err)
			return
		}
	}

	reqWithPhotos := r.URL.Query().Get("with_photos")
	if len(reqWithPhotos) != 0 {
		withPhotos, err = strconv.ParseBool(reqWithPhotos)
		if err != nil {
			httpUtils.WriteErrorResponse(w, http.StatusBadRequest, err)
			return
		}
	}

	products, err := h.ProductsService.GetProductsWithExtra(
		ctx, categoryUid, 0, 0, time.Time{}, time.Time{}, limit, 0, false, withPhotos, []uuid.UUID{},
	)
	if err != nil {
		httpUtils.WriteErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	resp := make([]httpEntity.ProductWithExtra, 0, len(products))
	for _, product := range products {
		resp = append(resp, httpEntity.ProductWithExtra{
			Product: httpEntity.ProductFromEntity(product.Product),
			Photos:  httpEntity.ProductPhotosFromEntity(product.Photos),
		})
	}

	httpUtils.WriteResponseJson(w, resp)
}
