package recomendationsSubrouter

import (
	"context"
	"net/http"
	"time"

	uuid "github.com/satori/go.uuid"
	httpEntity "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/entity"
	httpUtils "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/utils"
)

func (h *RecomendationsSubrouter) getRecomendations(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	// userUid := uuid.FromStringOrNil(r.URL.Query().Get("user_uid"))
	categoryUid := uuid.FromStringOrNil(r.URL.Query().Get("category_uid"))

	products, err := h.ProductsService.GetProducts(
		ctx, categoryUid, 0, 0, time.Time{}, time.Time{}, 50, 0,
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
