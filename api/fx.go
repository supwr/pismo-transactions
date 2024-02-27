package main

import (
	"github.com/supwr/pismo-transactions/api/handler"
	"github.com/supwr/pismo-transactions/config"
	"github.com/supwr/pismo-transactions/infrastructure/database"
	accountrepository "github.com/supwr/pismo-transactions/infrastructure/repository"
	"github.com/supwr/pismo-transactions/usecase/account"
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
			newApi,
			newAccountHandler,
			newAccountService,

			// repositories
			fx.Annotate(
				accountrepository.NewAccountAccountRepository,
				fx.As(new(account.RepositoryInterface)),
			),
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

func newConnection(cfg config.Config) (*gorm.DB, error) {
	return database.NewConnection(cfg)
}

func newAccountHandler(s *account.Service, l *slog.Logger) *handler.AccountHandler {
	return handler.NewAccountHandler(s, l)
}

func newApi(args ApiArgs) *Api {
	return NewApi(args)
}

func newAccountService(r account.RepositoryInterface) *account.Service {
	return account.NewService(r)
}
