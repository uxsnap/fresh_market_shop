package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/uxsnap/fresh_market_shop/backend/internal/app"
	"github.com/uxsnap/fresh_market_shop/backend/internal/config"
)

const defaultEnvPath = ".env"

func main() {
	if err := config.Init(defaultEnvPath); err != nil {
		log.Fatal(err)
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	app.New().Run(ctx)
}
