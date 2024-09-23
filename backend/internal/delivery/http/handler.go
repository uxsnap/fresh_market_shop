package deliveryHttp

import (
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type Config interface {
}

type Handler struct {
	router *chi.Mux
	config Config
}

func New(cfg Config) *Handler {
	h := &Handler{
		router: chi.NewRouter(),
		config: cfg,
	}

	h.router.Use(
		middleware.RequestID,
		middleware.RealIP,
		middleware.Logger,
		middleware.Recoverer,
		middleware.Timeout(60*time.Second),
	)

	h.router.Route("/health", h.HealthSubrouter)
	h.router.Route("/category", h.CategoriesSubrouter)

	return h
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.router.ServeHTTP(w, r)
}
