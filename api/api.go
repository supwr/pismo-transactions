package main

import (
	"github.com/gin-gonic/gin"
	"github.com/supwr/pismo-transactions/api/handler"
	"go.uber.org/fx"
	"log/slog"
)

type ApiArgs struct {
	fx.In

	Logger         *slog.Logger
	AccountHandler *handler.AccountHandler
}

type Api struct {
	logger         *slog.Logger
	accountHandler *handler.AccountHandler
}

func NewApi(a ApiArgs) *Api {
	return &Api{
		logger:         a.Logger,
		accountHandler: a.AccountHandler,
	}
}

func (a *Api) Serve() {
	app := gin.Default()

	// routes
	app.GET("/accounts/:accountId", a.accountHandler.GetAccountById)

	app.Run()
}
