package router

import (
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/jackc/pgx/v5/pgxpool"
)

func initAllRoutes(r *chi.Mux) {
	r.Route("/category", initCategoryRoute)
}

func New(db *pgxpool.Pool) *chi.Mux {
	r := chi.NewRouter()

	r.Use(
		middleware.RequestID,
		middleware.RealIP,
		middleware.Logger,
		middleware.Recoverer,
		middleware.WithValue("DB", db),
	)

	r.Use(middleware.Timeout(60 * time.Second))

	initAllRoutes(r)

	return r
}
