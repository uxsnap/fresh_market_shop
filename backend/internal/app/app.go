package app

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type App struct {
	router http.Handler
}

func New() *App {
	r := chi.NewRouter()

	r.Use(middleware.Timeout(60 * time.Second))

	return &App{
		router: r,
	}
}

func (a *App) Start(ctx context.Context) error {
	ch := make(chan error, 1)

	server := http.Server{
		Addr:    ":8000",
		Handler: a.router,
	}

	go func() {
		fmt.Println("Server is listening")

		err := server.ListenAndServe()

		if err != nil {
			ch <- fmt.Errorf("cannot start http server! %w", err)
		}

		close(ch)
	}()

	select {
	case err := <-ch:
		return err
	case <-ctx.Done():
		fmt.Println("\n === Server is shutting down. === ")

		timeout, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()

		return server.Shutdown(timeout)
	}

	return nil
}
