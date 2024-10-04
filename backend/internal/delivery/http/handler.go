package deliveryHttp

import (
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/subrouters"
	authSubrouter "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/subrouters/auth"
	categorySubrouter "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/subrouters/category"
	healthSubrouter "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/subrouters/health"
)

type Config interface {
}

type Handler struct {
	router *chi.Mux
	config Config

	deps subrouters.SubrouterDeps
}

func New(
	cfg Config,
	authService subrouters.AuthService,
	productsService subrouters.ProductsService,
) *Handler {
	h := &Handler{
		router: chi.NewRouter(),
		config: cfg,
		deps: subrouters.SubrouterDeps{
			AuthService:       authService,
			ProductsService:   productsService,
		},
	}

	h.router.Use(
		middleware.RequestID,
		middleware.RealIP,
		middleware.Logger,
		middleware.Recoverer,
		middleware.Timeout(60*time.Second),
	)

	h.router.Route("/health", healthSubrouter.New(h.deps))
	h.router.Route("/category", categorySubrouter.New(h.deps))
	h.router.Route("/auth", authSubrouter.New(h.deps))

	return h
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.router.ServeHTTP(w, r)
}
