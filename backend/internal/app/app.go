package app

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/uxsnap/fresh_market_shop/internal/db"
	"github.com/uxsnap/fresh_market_shop/internal/router"
)

type App struct {
	router http.Handler
}

func New() (*App, error) {
	db, err := db.New()

	if err != nil {
		return nil, err
	}

	return &App{
		router: router.New(db),
	}, nil
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
}
