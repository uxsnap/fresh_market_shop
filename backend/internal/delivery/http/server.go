package deliveryHttp

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type ServerConfig interface {
	ServiceHost() string
	ServicePort() string
}

type Server struct {
	router *chi.Mux
	config ServerConfig
}

func New(cfg ServerConfig) *Server {
	srv := &Server{
		router: chi.NewRouter(),
		config: cfg,
	}

	srv.router.Use(
		middleware.RequestID,
		middleware.RealIP,
		middleware.Logger,
		middleware.Recoverer,
		middleware.Timeout(60*time.Second),
	)

	srv.router.Route("/category", srv.CategoriesSubrouter)

	return srv
}

func (s *Server) Run(ctx context.Context) {
	ch := make(chan error, 1)

	server := http.Server{
		Addr:    s.config.ServiceHost() + ":" + s.config.ServicePort(),
		Handler: s.router,
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
