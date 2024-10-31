package recipesSubrouter

import (
	"context"
	"net/http"

	"github.com/go-chi/chi"
	uuid "github.com/satori/go.uuid"
	httpEntity "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/entity"
	httpUtils "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/utils"
)

func (h *RecipesSubrouter) GetRecipeSteps(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	uid, err := uuid.FromString(chi.URLParam(r, "recipe_uid"))
	if err != nil {
		httpUtils.WriteErrorResponse(w, http.StatusBadRequest, nil)
		return
	}

	steps, err := h.RecipesService.GetRecipeSteps(ctx, uid)

	if err != nil {
		httpUtils.WriteErrorResponse(w, http.StatusInternalServerError, nil)
		return
	}

	resp := make([]httpEntity.RecipeStep, 0, len(steps))
	for _, step := range steps {
		resp = append(resp, httpEntity.RecipeStepFromEntity(step))
	}

	httpUtils.WriteResponseJson(w, resp)
}
