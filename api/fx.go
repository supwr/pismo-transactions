package main

import (
	"github.com/supwr/pismo-transactions/api/handler"
	"github.com/supwr/pismo-transactions/internal/account"
	"github.com/supwr/pismo-transactions/internal/transaction"
	"github.com/supwr/pismo-transactions/pkg/clock"
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
			newClock,

			//handlers
			newAccountHandler,
			newTransactionHandler,

			//services
			newAccountService,
			newTransactionService,

			// repositories
			fx.Annotate(
				account.NewRepository,
				fx.As(new(account.RepositoryInterface)),
			),
			fx.Annotate(
				transaction.NewRepository,
				fx.As(new(transaction.RepositoryInterface)),
			),
		),
	}

	return fx.New(append(options, o...)...)
}

func newLogger() *slog.Logger {
	return slog.New(slog.NewTextHandler(os.Stderr, nil))
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

func newTransactionService(r transaction.RepositoryInterface, a *account.Service, c clock.Clock) *transaction.Service {
	return transaction.NewService(r, a, c)
}

func newClock() clock.Clock {
	return clock.NewClock()
}
