package supportSubrouterTickets

import (
	"github.com/go-chi/chi"
	"github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/subrouters"
)

type SupportSubrouterTickets struct {
	subrouters.SubrouterDeps
}

func New(deps subrouters.SubrouterDeps) func(r chi.Router) {
	st := SupportSubrouterTickets{deps}

	return func(r chi.Router) {
		r.Get("/{uid}", st.GetTicketByUid)
		r.Get("/", st.GetTickets)

		r.Group(func(r chi.Router) {
			r.Use(st.Middleware.AuthOrPass)

			r.Post("/", st.CreateTicket)
		})

		r.Group(func(r chi.Router) {
			r.Use(st.Middleware.Auth)

			r.Put("/", st.EditTicket)
			r.Post("/take/{uid}", st.TakeTicket)
		})
	}
}
