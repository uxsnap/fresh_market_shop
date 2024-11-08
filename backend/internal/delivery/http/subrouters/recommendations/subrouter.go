package recommendationsSubrouter

import (
	"github.com/go-chi/chi"
	"github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/subrouters"
)

type RecommendationsSubrouter struct {
	subrouters.SubrouterDeps
}

func New(deps subrouters.SubrouterDeps) func(r chi.Router) {
	rs := RecommendationsSubrouter{deps}

	return func(r chi.Router) {

		r.Group(func(r chi.Router) {
			r.Use(rs.Middleware.AuthOrPass)

			r.Get("/", rs.getRecommendations)
		})
	}
}
