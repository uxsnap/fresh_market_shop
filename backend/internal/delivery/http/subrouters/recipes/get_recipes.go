package recipesSubrouter

import (
	"context"
	"net/http"

	httpEntity "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/entity"
	httpUtils "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/utils"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
	errorWrapper "github.com/uxsnap/fresh_market_shop/backend/internal/error_wrapper"
)

func (h *RecipesSubrouter) GetRecipes(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	qFilters, err := entity.NewQueryFiltersParser().
		WithRequired(
			entity.QueryFieldPage,
			entity.QueryFieldLimit,
		).
		WithAllowed(
			entity.QueryFieldName,
		).
		ParseQuery(r.URL.Query())
	if err != nil {
		httpUtils.WriteErrorResponse(w, http.StatusBadRequest, errorWrapper.NewError("parsing query params error", err.Error()))
		return
	}

	recipes, err := h.RecipesService.GetRecipes(ctx, qFilters)

	if err != nil {
		httpUtils.WriteErrorResponse(w, http.StatusBadRequest, errorWrapper.NewError(
			errorWrapper.InternalError, err.Error(),
		))
		return
	}

	resp := httpEntity.RecipesWithTotal{}

	httpRecipes := make([]httpEntity.Recipe, 0, len(recipes.Recipes))

	for _, recipe := range recipes.Recipes {
		httpRecipes = append(httpRecipes, httpEntity.RecipeFromEntity(recipe))
	}

	resp.Recipes = httpRecipes
	resp.Total = recipes.Total

	httpUtils.WriteResponseJson(w, resp)
}
