package main

import (
	"github.com/supwr/pismo-transactions/api/handler"
	"github.com/supwr/pismo-transactions/config"
	"github.com/supwr/pismo-transactions/infrastructure/database"
	"github.com/supwr/pismo-transactions/infrastructure/repository"
	"github.com/supwr/pismo-transactions/usecase/account"
	"github.com/supwr/pismo-transactions/usecase/operation_type"
	"github.com/supwr/pismo-transactions/usecase/transaction"
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

func newTransactionService(r transaction.RepositoryInterface, o *operation_type.Service, a *account.Service) *transaction.Service {
	return transaction.NewService(r, o, a)
}
