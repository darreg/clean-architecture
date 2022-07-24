package main

import (
	"fmt"

	"github.com/alrund/yp-1-project/internal/application/app"
	"github.com/alrund/yp-1-project/internal/infrastructure/adapter"
	"github.com/alrund/yp-1-project/internal/infrastructure/handler"
	"github.com/alrund/yp-1-project/internal/infrastructure/service"
)

func main() {
	var (
		logger = adapter.NewLogger()
		router = adapter.NewRouter()
		flags  = service.GetFlags()
		config = service.GetConfig(flags, logger)
	)

	fmt.Println(config.RunAddress)

	a := app.NewApp(
		config,
		logger,
		router,
	)

	r := a.Router.WithPrefix("/api/user")
	r.Post("/register", handler.RegisterHandler(a))
	r.Post("/login", handler.LoginHandler(a))

	if err := a.Run(); err != nil {
		logger.Fatal(err)
	}
}
