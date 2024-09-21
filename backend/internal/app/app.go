package app

import (
	"context"

	"github.com/uxsnap/fresh_market_shop/internal/config"
	deliveryHttp "github.com/uxsnap/fresh_market_shop/internal/delivery/http"
)

type App struct {
	httpServer *deliveryHttp.Server
}

func New() *App {
	// убрать внутрь service provider
	// _, _ = db.New()

	return &App{
		httpServer: deliveryHttp.New(config.NewHttpConfig()),
	}
}

func (a *App) Run(ctx context.Context) {
	a.httpServer.Run(ctx)
	<-ctx.Done()
}
