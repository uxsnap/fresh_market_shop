package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/uxsnap/fresh_market_shop/internal/app"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	app, appErr := app.New()

	if appErr != nil {
		log.Fatal(appErr)
	}

	err = app.Start(ctx)

	if err != nil {
		log.Fatal(err)
	}
}
