package recipesSubrouter

import (
	"context"
	"net/http"

	"github.com/go-chi/chi"
	uuid "github.com/satori/go.uuid"
	httpEntity "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/entity"
	httpUtils "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/utils"
)

func (h *RecipesSubrouter) GetRecipeProducts(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	uid, err := uuid.FromString(chi.URLParam(r, "recipe_uid"))
	if err != nil {
		httpUtils.WriteErrorResponse(w, http.StatusBadRequest, nil)
		return
	}

	products, err := h.RecipesService.GetRecipesProducts(ctx, uid)

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
}
