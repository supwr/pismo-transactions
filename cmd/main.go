package main

import (
	"github.com/shopspring/decimal"
	"github.com/supwr/pismo-transactions/internal/config"
	"github.com/supwr/pismo-transactions/internal/infrastructure/database"
	"go.uber.org/fx"
)

const devEnv = "DEV"

func main() {
	decimal.MarshalJSONWithoutQuotes = true

	app := createApp(
		fx.Invoke(func(cfg config.Config, migration *database.Migration) {
			if cfg.Environment == devEnv {
				migration.CreateSchema()
				migration.Migrate()
			}
		}),
		fx.Invoke(func(s fx.Shutdowner) { _ = s.Shutdown() }),
	)

	app.Run()
}
