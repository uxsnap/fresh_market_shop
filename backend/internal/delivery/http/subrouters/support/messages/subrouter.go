package supportSubrouterMessages

import (
	"github.com/go-chi/chi"
	"github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/subrouters"
)

type SupportSubrouterMessages struct {
	subrouters.SubrouterDeps
}

func New(deps subrouters.SubrouterDeps) func(r chi.Router) {
	st := SupportSubrouterMessages{deps}

	return func(r chi.Router) {
		r.Use(st.Middleware.Auth)

		r.Post("/", st.AddTicketMessage)
		r.Put("/", st.EditTicketMessage)
		r.Get("/", st.GetTicketMessages)
	}
}
