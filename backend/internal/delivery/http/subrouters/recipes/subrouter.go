package recipesSubrouter

import (
	"github.com/go-chi/chi"
	"github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/subrouters"
)

type RecipesSubrouter struct {
	subrouters.SubrouterDeps
}

func New(deps subrouters.SubrouterDeps) func(r chi.Router) {
	rs := RecipesSubrouter{deps}

	return func(r chi.Router) {
		r.Get("/{recipe_uid}", rs.GetRecipeByUid)
		r.Get("/", rs.GetRecipes)
		r.Get("/{recipe_uid}/products", rs.GetRecipeProducts)
		r.Get("/{recipe_uid}/steps", rs.GetRecipeSteps)

		r.Group(func(r chi.Router) {
			r.Use(rs.Middleware.Auth)

			r.Post("/", rs.CreateRecipe)

			r.Put("/", rs.UpdateRecipe)

			r.Delete("/{recipe_uid}", rs.DeleteRecipe)
		})
	}

}
