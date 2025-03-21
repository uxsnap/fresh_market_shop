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

		r.Group(func(r chi.Router) {
			r.Use(us.Middleware.Auth)

			r.Get("/{user_uid}", us.getUser)
			r.Post("/{user_uid}/photo", us.uploadPhoto)
			r.Put("/", us.updateUser)

			r.Post("/{user_uid}/delivery_address", us.addDeliveryAddress)
			r.Put("/{user_uid}/delivery_address", us.updateDeliveryAddress)
			r.Delete("/{user_uid}/delivery_address/{address_uid}", us.deleteDeliveryAddress)

			r.Get("/{user_uid}/delivery_address", us.getUserDeliveryAddresses)
			r.Get("/delivery_address/{address_uid}", us.getDeliveryAddress)
		})
	}
}
