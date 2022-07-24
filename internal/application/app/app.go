package app

import "github.com/alrund/yp-1-project/internal/domain/port"

type App struct {
	Config          *Config
	Logger          port.Logger
	Router          port.Router
	UserRepository  port.UserRepository
	OrderRepository port.OrderRepository
}

func NewApp(
	config *Config,
	logger port.Logger,
	router port.Router,
	userRepository port.UserRepository,
	orderRepository port.OrderRepository,
) *App {
	return &App{
		Config:          config,
		Logger:          logger,
		Router:          router,
		UserRepository:  userRepository,
		OrderRepository: orderRepository,
	}
}

func (a *App) Run() error {
	return a.Serve()
}

func (a *App) Stop() error {
	return nil
}
