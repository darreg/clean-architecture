package app

import "github.com/alrund/yp-1-project/internal/domain/port"

type App struct {
	Config             *Config
	Logger             port.Logger
	Router             port.Router
	Encryptor          port.Encryptor
	Hasher             port.PasswordHasher
	Cooker             port.Cooker
	Storage            port.Storage
	Transactor         port.Transactor
	UserRepository     port.UserRepository
	OrderRepository    port.OrderRepository
	WithdrawRepository port.WithdrawRepository
}

func NewApp(
	config *Config,
	logger port.Logger,
	router port.Router,
	encryptor port.Encryptor,
	hasher port.PasswordHasher,
	cooker port.Cooker,
	storage port.Storage,
	transactor port.Transactor,
	userRepository port.UserRepository,
	orderRepository port.OrderRepository,
	withRepository port.WithdrawRepository,
) *App {
	return &App{
		Config:             config,
		Logger:             logger,
		Router:             router,
		Encryptor:          encryptor,
		Hasher:             hasher,
		Cooker:             cooker,
		Storage:            storage,
		Transactor:         transactor,
		UserRepository:     userRepository,
		OrderRepository:    orderRepository,
		WithdrawRepository: withRepository,
	}
}

func (a *App) Run() error {
	return a.Serve()
}

func (a *App) Stop() error {
	return nil
}
