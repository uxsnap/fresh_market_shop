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

		r.Group(func(r chi.Router) {
			r.Use(s.Middleware.Auth)

			r.Get("/{user_uid}/{order_uid}", s.GetOrder)
			r.Get("/{user_uid}/history", s.GetHistory)
			r.Post("/", s.CreateOrder)
			r.Post("/pay", s.PayOrder)
			r.Get("/products", s.GetOrderedProducts)
		})
	}
}
