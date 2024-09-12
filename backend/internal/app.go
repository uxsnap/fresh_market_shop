package internal

import (
	"net/http"

	"github.com/go-chi/chi"
)

type App struct {
	router http.Handler
}

func New() *App {
	r := chi.NewRouter()

	return &App{
		router: r,
	}
}

func (a *App) Start() {
	http.ListenAndServe(":8000", a.router)
}
