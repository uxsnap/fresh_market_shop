package recomendationsSubrouter

import (
	"github.com/go-chi/chi"
	"github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/subrouters"
)

type RecomendationsSubrouter struct {
	subrouters.SubrouterDeps
}

func New(deps subrouters.SubrouterDeps) func(r chi.Router) {
	rs := RecomendationsSubrouter{deps}

	return func(r chi.Router) {
		r.Get("/", rs.getRecomendations)
	}
}
