package main

import (
	"github.com/alrund/yp-1-project/internal/application/app"
	"github.com/alrund/yp-1-project/internal/infrastructure/adapter"
	"github.com/alrund/yp-1-project/internal/infrastructure/builder"
	"github.com/alrund/yp-1-project/internal/infrastructure/handler"
	"github.com/alrund/yp-1-project/internal/infrastructure/middleware"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	logger := adapter.NewLogger()

	config, err := app.NewConfig(adapter.NewConfigLoader())
	if err != nil {
		logger.Fatal(err)
	}

	a, err := builder.Builder(config, logger)
	if err != nil {
		logger.Fatal(err)
	}

	initRoutes(a)

	if err := a.Run(); err != nil {
		logger.Fatal(err)
	}
}

func initRoutes(a *app.App) {
	r := a.Router.WithPrefix("/api/user")
	r.Post("/register", handler.RegisterHandler(a))
	r.Post("/login", handler.LoginHandler(a))

	r.Post("/orders", handler.AddOrderHandler(a))
	r.Get("/orders", handler.OrderListHandler(a))

	r.Get("/withdrawals", handler.WithdrawListHandler(a))
	r.Post("/balance/withdraw", handler.AddWithdrawHandler(a))
	r.Get("/balance", handler.BalanceHandler(a))

	r.Use(middleware.RequestLog(a))
	r.Use(middleware.Auth(a))
}
