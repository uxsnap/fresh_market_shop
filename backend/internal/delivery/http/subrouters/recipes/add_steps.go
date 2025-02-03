package recipesSubrouter

import (
	"context"
	"net/http"

	httpEntity "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/entity"
	httpUtils "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/utils"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
	errorWrapper "github.com/uxsnap/fresh_market_shop/backend/internal/error_wrapper"
)

func (h *RecipesSubrouter) AddSteps(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	userInfo, err := httpEntity.AuthUserInfoFromContext(r.Context())
	if err != nil {
		httpUtils.WriteErrorResponse(w, http.StatusBadRequest, errorWrapper.NewError(
			errorWrapper.JsonParsingError, "не удалось найти юзера",
		))
		return
	}

	if userInfo.Role != "admin" {
		httpUtils.WriteErrorResponse(w, http.StatusBadRequest, errorWrapper.NewError(
			errorWrapper.JsonParsingError, "нет разрешений на создание рецепта",
		))
		return
	}

	recipeSteps := []httpEntity.RecipeStep{}
	if err := httpUtils.DecodeJsonRequest(r, recipeSteps); err != nil {
		httpUtils.WriteErrorResponse(w, http.StatusBadRequest, errorWrapper.NewError(
			errorWrapper.JsonParsingError, "не удалось распарсить тело запроса",
		))
		return
	}

	rSteps := make([]entity.RecipeStep, len(recipeSteps))

	for ind, r := range recipeSteps {
		rSteps[ind] = httpEntity.RecipeStepToEntity(r)
	}

	err = h.RecipesService.AddRecipeSteps(ctx, rSteps)
	if err != nil {
		httpUtils.WriteErrorResponse(w, http.StatusInternalServerError, errorWrapper.NewError(
			errorWrapper.InternalError, err.Error(),
		))
		return
	}

	httpUtils.WriteResponseJson(w, httpEntity.UUID{})
}
