package usersSubrouter

import (
	"github.com/go-chi/chi"
	"github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/subrouters"
)

type UsersSubrouter struct {
	subrouters.SubrouterDeps
}

func New(deps subrouters.SubrouterDeps) func(r chi.Router) {
	us := UsersSubrouter{deps}

	return func(r chi.Router) {
		r.Post("/create", us.createUser)
		r.Post("/update", us.updateUser)
		r.Get("/{user_uid}", us.getUser)

		r.Post("/{user_uid}/delivery_address/create", us.addDeliveryAddress)
		r.Post("/{user_uid}/delivery_address/update", us.updateDeliveryAddress)
		r.Post("/delivery_address/{address_uid}/delete", us.deleteDeliveryAddress)

		r.Get("/{user_uid}/delivery_addresses", us.getUserDeliveryAddresses)
		r.Get("/delivery_address/{address_uid}", us.getDeliveryAddress)
	}
}
