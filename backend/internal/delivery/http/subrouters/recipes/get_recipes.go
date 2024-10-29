package recipesSubrouter

import (
	"context"
	"net/http"

	httpUtils "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/utils"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
	errorWrapper "github.com/uxsnap/fresh_market_shop/backend/internal/error_wrapper"
)

func (h *RecipesSubrouter) GetRecipes(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	recipes, err := h.RecipesService.GetRecipes(ctx, entity.QueryFilters{})

	if err != nil {
		httpUtils.WriteErrorResponse(w, http.StatusBadRequest, errorWrapper.NewError(
			errorWrapper.JsonParsingError, "не удалось распарсить тело запроса",
		))
		return
	}

	// cookingTime int64

	httpUtils.WriteResponseJson(w, recipes)
}
