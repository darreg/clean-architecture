package main

import (
	"errors"

	"github.com/alrund/yp-1-project/internal/application/app"
	"github.com/alrund/yp-1-project/internal/domain/port"
	"github.com/alrund/yp-1-project/internal/infrastructure/adapter"
	"github.com/alrund/yp-1-project/internal/infrastructure/handler"
	"github.com/alrund/yp-1-project/internal/infrastructure/middleware"
	"github.com/alrund/yp-1-project/internal/infrastructure/service"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
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
	config, err := app.NewConfig()
	if err != nil {
		return nil, err
	}

	storage, err := service.NewStorage(config.DatabaseURI)
	if err != nil {
		return nil, err
	}

	err = migrations(storage)
	if err != nil {
		return nil, err
	}

	var (
		encryption      = service.NewEncryption(config.CipherPass)
		hasher          = service.NewPasswordHasher()
		router          = adapter.NewRouter()
		cooker          = service.NewCookies()
		userRepository  = adapter.NewUserRepository(storage.Connect, hasher)
		orderRepository = adapter.NewOrderRepository(storage.Connect)
	)

	return app.NewApp(
		config,
		logger,
		router,
		encryption,
		hasher,
		cooker,
		userRepository,
		orderRepository,
	), nil
}

func migrations(storage *service.Storage) error {
	driver, err := postgres.WithInstance(storage.Connect, &postgres.Config{})
	if err != nil {
		return err
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"postgres", driver)
	if err != nil {
		return err
	}

	err = m.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return err
	}

	return nil
}
