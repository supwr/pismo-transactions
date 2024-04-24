package main

import (
	"github.com/supwr/pismo-transactions/pkg/database"
	"go.uber.org/fx"
	"log/slog"
	"os"
)

func createApp(o ...fx.Option) *fx.App {
	options := []fx.Option{
		database.Module(),
		fx.Provide(
			newLogger,
		),
	}

	return fx.New(append(options, o...)...)
}

func newLogger() *slog.Logger {
	return slog.New(slog.NewTextHandler(os.Stderr, nil))
}
