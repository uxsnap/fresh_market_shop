package deliveryHttp

import (
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	cors "github.com/go-chi/cors"
	"github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/subrouters"
	assetsSubrouter "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/subrouters/assets"
	authSubrouter "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/subrouters/auth"
	categoriesSubrouter "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/subrouters/categories"
	healthSubrouter "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/subrouters/health"
	ordersSubrouter "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/subrouters/orders"
	productsSubrouter "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/subrouters/products"
	recommendationsSubrouter "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/subrouters/recommendations"
	searchSubrouter "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/subrouters/search"
	usersSubrouter "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/subrouters/users"
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
	usersService subrouters.UsersService,
	recipesService subrouters.RecipesService,
	ordersService subrouters.OrdersService,
) *Handler {
	h := &Handler{
		router: chi.NewRouter(),
		config: cfg,
		deps: subrouters.SubrouterDeps{
			AuthService:     authService,
			ProductsService: productsService,
			UsersService:    usersService,
			RecipesService:  recipesService,
			OrdersService:   ordersService,
		},
	}

	h.router.Use(
		middleware.RequestID,
		middleware.RealIP,
		middleware.Logger,
		middleware.Recoverer,
		middleware.Timeout(60*time.Second),
	)

	h.router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	h.router.Route("/auth", authSubrouter.New(h.deps))
	h.router.Route("/health", healthSubrouter.New(h.deps))
	h.router.Route("/categories", categoriesSubrouter.New(h.deps))
	h.router.Route("/products", productsSubrouter.New(h.deps))
	h.router.Route("/users", usersSubrouter.New(h.deps))
	h.router.Route("/search", searchSubrouter.New(h.deps))
	h.router.Route("/assets", assetsSubrouter.New(h.deps))
	h.router.Route("/recommendations", recommendationsSubrouter.New(h.deps))
	h.router.Route("/orders", ordersSubrouter.New(h.deps))

	return h
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.router.ServeHTTP(w, r)
}
