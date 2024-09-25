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

	authService       AuthService
	productsService   ProductsService
	categoriesService CategoriesService
}

func New(
	cfg Config,
	authService AuthService,
	productsService ProductsService,
	categoriesService CategoriesService,
) *Handler {
	h := &Handler{
		router:            chi.NewRouter(),
		config:            cfg,
		authService:       authService,
		productsService:   productsService,
		categoriesService: categoriesService,
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
	h.router.Route("/auth", h.AuthSubrouter)

	return h
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.router.ServeHTTP(w, r)
}
