package deliverySubrouter

import (
	"github.com/go-chi/chi"
	"github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/subrouters"
)

type DeliverySubrouter struct {
	subrouters.SubrouterDeps
}

func New(deps subrouters.SubrouterDeps) func(r chi.Router) {
	ds := DeliverySubrouter{deps}

	return func(r chi.Router) {
		r.Use(ds.Middleware.Auth)

		r.Put("/", ds.UpdateDelivery)
		r.Get("/{uid}", ds.GetDeliveryByUid)
		r.Get("/by_order/{order_uid}", ds.GetDeliveryByOrderUid)
	}
}
