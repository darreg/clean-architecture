package app

import "github.com/alrund/yp-1-project/internal/domain/port"

type App struct {
	Config          *Config
	Logger          port.Logger
	Router          port.Router
	Encryptor       port.Encryptor
	Hasher          port.PasswordHasher
	Cooker          port.Cooker
	UserRepository  port.UserRepository
	OrderRepository port.OrderRepository
}

func NewApp(
	config *Config,
	logger port.Logger,
	router port.Router,
	encryptor port.Encryptor,
	hasher port.PasswordHasher,
	cooker port.Cooker,
	userRepository port.UserRepository,
	orderRepository port.OrderRepository,
) *App {
	return &App{
		Config:          config,
		Logger:          logger,
		Router:          router,
		Encryptor:       encryptor,
		Hasher:          hasher,
		Cooker:          cooker,
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
