package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/uxsnap/fresh_market_shop/internal/app"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	app := app.New()

	err := app.Start(ctx)

	if err != nil {
		fmt.Println(err)
	}
}
