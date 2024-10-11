package searchSubrouter

import (
	"github.com/go-chi/chi"
	"github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/subrouters"
)

type SearchSubrouter struct {
	subrouters.SubrouterDeps
}

func New(deps subrouters.SubrouterDeps) func(r chi.Router) {
	s := SearchSubrouter{deps}

	return func(r chi.Router) {
		r.Get("/products", s.searchProducts)
		r.Get("/categories", s.searchCategories)
		r.Get("/", s.multipleSearch)
	}
}
