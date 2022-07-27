package main

import (
	"github.com/alrund/yp-1-project/internal/application/app"
	"github.com/alrund/yp-1-project/internal/domain/port"
	"github.com/alrund/yp-1-project/internal/infrastructure/adapter"
	"github.com/alrund/yp-1-project/internal/infrastructure/handler"
	"github.com/alrund/yp-1-project/internal/infrastructure/middleware"
	"github.com/alrund/yp-1-project/internal/infrastructure/repository"
	"github.com/alrund/yp-1-project/internal/infrastructure/service"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	logger := adapter.NewLogger()

	a, err := builder(logger)
	if err != nil {
		logger.Fatal(err)
	}

	routes(a)

	if err := a.Run(); err != nil {
		logger.Fatal(err)
	}
}

func routes(a *app.App) {
	r := a.Router.WithPrefix("/api/user")
	r.Post("/register", handler.RegisterHandler(a))
	r.Post("/login", handler.LoginHandler(a))

	r.Use(middleware.Auth(a))
}

func builder(logger port.Logger) (*app.App, error) {
	config, err := app.NewConfig(adapter.NewConfigLoader())
	if err != nil {
		return nil, err
	}

	storage, err := service.NewStorage(config.DatabaseURI)
	if err != nil {
		return nil, err
	}

	var (
		router          = adapter.NewRouter()
		cooker          = adapter.NewCooker()
		hasher          = adapter.NewHasher()
		encryptor       = adapter.NewEncryptor(config.CipherPass)
		userRepository  = repository.NewUserRepository(storage.Connect, hasher)
		orderRepository = repository.NewOrderRepository(storage.Connect)
	)

	return app.NewApp(
		config,
		logger,
		router,
		encryptor,
		hasher,
		cooker,
		userRepository,
		orderRepository,
	), nil
}
