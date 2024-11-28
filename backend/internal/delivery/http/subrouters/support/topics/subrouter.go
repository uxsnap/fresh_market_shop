package supportSubrouterTopics

import (
	"github.com/go-chi/chi"
	"github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/subrouters"
)

type SupportSubrouterTopics struct {
	subrouters.SubrouterDeps
}

func New(deps subrouters.SubrouterDeps) func(r chi.Router) {
	st := SupportSubrouterTopics{deps}

	return func(r chi.Router) {
		r.Get("/{uid}", st.GetTicketsTopicByUid)
		r.Get("/", st.GetAllTicketsTopics)
		r.Get("/{uid}/solutions", st.GetTopicSolutions)

		r.Group(func(r chi.Router) {
			r.Use(st.Middleware.Auth)

			r.Post("/", st.CreateTicketsTopic)
			r.Put("/", st.UpdateTicketsTopic)
		})

	}

}
