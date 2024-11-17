package paymentsSubrouter

import (
	"github.com/go-chi/chi"
	"github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/subrouters"
)

type PaymentsSubrouter struct {
	subrouters.SubrouterDeps
}

func New(deps subrouters.SubrouterDeps) func(r chi.Router) {
	rs := PaymentsSubrouter{deps}

	return func(r chi.Router) {

		r.Group(func(r chi.Router) {
			r.Use(rs.Middleware.Auth)

			r.Post("/cards", rs.AddUserPaymentCard)
			r.Get("/cards/{card_uid}", rs.GetPaymentCard)
			r.Get("/cards/by_user/{user_uid}", rs.GetUserPaymentCards)

			r.Delete("/cards/{card_uid}", rs.DeletePaymentCard)
			r.Delete("/cards/by_user/{user_uid}", rs.DeleteUserPaymentCards)

			r.Post("/", rs.CreatePayment)
			r.Get("/{payment_uid}", rs.GetPayment)
			r.Get("/by_order/{order_uid}", rs.GetOrderPayment)
			r.Get("/", rs.GetPayments)

			// TODO: remove

		})
	}
}
