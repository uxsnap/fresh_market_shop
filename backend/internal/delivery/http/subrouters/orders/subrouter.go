package ordersSubrouter

import (
	"github.com/go-chi/chi"
	"github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/subrouters"
)

type OrdersSubrouter struct {
	subrouters.SubrouterDeps
}

func New(deps subrouters.SubrouterDeps) func(r chi.Router) {
	s := OrdersSubrouter{deps}

	return func(r chi.Router) {
		r.Get("/history/{user_uid}", s.getHistory)

		r.Group(func(r chi.Router) {
			r.Use(s.Middleware.Auth)

			r.Post("/", s.CreateOrder)
		})

	}
}
