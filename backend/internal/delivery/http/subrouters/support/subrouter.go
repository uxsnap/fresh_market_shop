package supportSubrouter

import (
	"github.com/go-chi/chi"
	"github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/subrouters"
	messagesSubrouter "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/subrouters/support/messages"
	solutionsSubrouter "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/subrouters/support/solutions"
	ticketsSubrouter "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/subrouters/support/tickets"
	topicsSubrouter "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/subrouters/support/topics"
)

func New(deps subrouters.SubrouterDeps) func(r chi.Router) {
	return func(r chi.Router) {
		r.Route("/tickets/topics", topicsSubrouter.New(deps))
		r.Route("/tickets", ticketsSubrouter.New(deps))
		r.Route("/tickets/{ticket_uid}/messages", messagesSubrouter.New(deps))
		r.Route("/tickets/{ticket_uid}/solution", solutionsSubrouter.New(deps))
	}
}
