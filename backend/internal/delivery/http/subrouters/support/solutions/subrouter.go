package supportSubrouterSolutions

import (
	"github.com/go-chi/chi"
	"github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/subrouters"
)

type SupportSubrouterSolutions struct {
	subrouters.SubrouterDeps
}

func New(deps subrouters.SubrouterDeps) func(r chi.Router) {
	st := SupportSubrouterSolutions{deps}

	return func(r chi.Router) {
		r.Use(st.Middleware.Auth)

		r.Post("/", st.CreateTicketSolution)
		r.Get("/", st.GetTicketSolution)
	}
}
