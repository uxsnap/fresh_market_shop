package main

import "github.com/uxsnap/fresh_market_shop/internal"

func main() {
	app := internal.New()

	app.Start()
}
