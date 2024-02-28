package main

import (
	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
	"github.com/supwr/pismo-transactions/api/handler"
	_ "github.com/supwr/pismo-transactions/docs"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"go.uber.org/fx"
)

// @title           Transactions API
// @version         1.0
func main() {
	decimal.MarshalJSONWithoutQuotes = true

	app := createApp(
		fx.Invoke(func(accountHandler *handler.AccountHandler) {
			api := gin.Default()

			// routes
			api.GET("/accounts/:accountId", accountHandler.GetAccountById)
			api.POST("/accounts", accountHandler.CreateAccount)
			api.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

			api.Run()
		}),
		fx.Invoke(func(s fx.Shutdowner) { _ = s.Shutdown() }),
	)

	app.Run()
}
