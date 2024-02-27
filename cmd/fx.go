package main

import (
	"github.com/supwr/pismo-transactions/config"
	"github.com/supwr/pismo-transactions/infrastructure/database"
	"go.uber.org/fx"
	"gorm.io/gorm"
	"log/slog"
	"os"
)

func createApp(o ...fx.Option) *fx.App {
	options := []fx.Option{
		fx.Provide(
			newConfig,
			newLogger,
			newConnection,
			newMigration,
		),
	}

	return fx.New(append(options, o...)...)
}

func newConfig() (config.Config, error) {
	return config.NewConfig()
}

func newLogger() *slog.Logger {
	return slog.New(slog.NewTextHandler(os.Stderr, nil))
}

func newMigration(db *gorm.DB, cfg config.Config, logger *slog.Logger) *database.Migration {
	return database.NewMigration(db, cfg, logger)
}

func newConnection(cfg config.Config) (*gorm.DB, error) {
	return database.NewConnection(
		cfg,
	)
}
