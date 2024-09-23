package app

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

type App struct {
	serviceProvider *serviceProvider
}

func New() *App {

	return &App{
		serviceProvider: newServiceProvider(),
	}
}

func (a *App) Run(ctx context.Context) {
	a.RunHTTPServer(ctx)
	<-ctx.Done()
}

func (a *App) RunHTTPServer(ctx context.Context) {
	ch := make(chan error, 1)

	addr := a.serviceProvider.ConfigExternalApi().AuthServiceGrpcHost() + ":" + a.serviceProvider.ConfigExternalApi().AuthServiceGrpcPort()

	server := http.Server{
		Addr:    addr,
		Handler: a.serviceProvider.HandlerHTTP(ctx),
	}

	go func() {
		fmt.Println("Server is listening")

		err := server.ListenAndServe()

		if err != nil {
			ch <- fmt.Errorf("cannot start http server! %w", err)
		}

		close(ch)
	}()

	go func() {
		select {
		case err := <-ch:
			log.Printf("http server canceled with error %v", err)
			return
		case <-ctx.Done():
			fmt.Println("\n === Server is shutting down. === ")

			timeout, cancel := context.WithTimeout(context.Background(), time.Second*10)
			defer cancel()

			if err := server.Shutdown(timeout); err != nil {
				log.Printf("http server shutdown error %v", err)
			}
			return
		}
	}()
}
