package database

import "go.uber.org/fx"

func Module() fx.Option {
	return fx.Module("database",
		fx.Provide(
			NewConfig,
			NewConnection,
			NewMigration,
		),
	)
}
