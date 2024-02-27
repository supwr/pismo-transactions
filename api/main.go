package main

import (
	"github.com/shopspring/decimal"
	"go.uber.org/fx"
)

func main() {
	decimal.MarshalJSONWithoutQuotes = true

	app := createApp(
		fx.Invoke(func(api *Api) {
			api.Serve()
		}),
		fx.Invoke(func(s fx.Shutdowner) { _ = s.Shutdown() }),
	)

	app.Run()
}
