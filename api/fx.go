package main

import (
	"github.com/supwr/pismo-transactions/api/handler"
	"github.com/supwr/pismo-transactions/internal/config"
	"github.com/supwr/pismo-transactions/internal/infrastructure/database"
	"github.com/supwr/pismo-transactions/internal/infrastructure/repository"
	"github.com/supwr/pismo-transactions/internal/usecase/account"
	"github.com/supwr/pismo-transactions/internal/usecase/operation_type"
	"github.com/supwr/pismo-transactions/internal/usecase/transaction"

	"github.com/supwr/pismo-transactions/pkg/clock"

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
			newClock,

			//handlers
			newAccountHandler,
			newTransactionHandler,

			//services
			newOperationTypeService,
			newAccountService,
			newTransactionService,

			// repositories
			fx.Annotate(
				repository.NewAccountRepository,
				fx.As(new(account.RepositoryInterface)),
			),
			fx.Annotate(
				repository.NewOperationTypeRepository,
				fx.As(new(operation_type.RepositoryInterface)),
			),
			fx.Annotate(
				repository.NewTransactionRepository,
				fx.As(new(transaction.RepositoryInterface)),
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

func newTransactionHandler(s *transaction.Service, l *slog.Logger) *handler.TransactionHandler {
	return handler.NewTransactionHandler(s, l)
}

func newAccountService(r account.RepositoryInterface) *account.Service {
	return account.NewService(r)
}

func newOperationTypeService(r operation_type.RepositoryInterface) *operation_type.Service {
	return operation_type.NewService(r)
}

func newTransactionService(r transaction.RepositoryInterface, o *operation_type.Service, a *account.Service, c clock.Clock) *transaction.Service {
	return transaction.NewService(r, o, a, c)
}

func newClock() clock.Clock {
	return clock.NewClock()
}
