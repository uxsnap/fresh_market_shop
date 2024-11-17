package addressesSubrouter

import (
	"github.com/go-chi/chi"
	"github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/subrouters"
)

type AddressesSubrouter struct {
	subrouters.SubrouterDeps
}

func New(deps subrouters.SubrouterDeps) func(r chi.Router) {
	as := AddressesSubrouter{deps}

	return func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Use(as.Middleware.Auth)

			r.Get("/cities", as.GetCities)
		})
	}
}
