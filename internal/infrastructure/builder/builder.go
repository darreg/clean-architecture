package builder

import (
	"github.com/alrund/yp-1-project/internal/application/app"
	"github.com/alrund/yp-1-project/internal/domain/port"
	"github.com/alrund/yp-1-project/internal/infrastructure/adapter"
	"github.com/alrund/yp-1-project/internal/infrastructure/repository"
)

func Builder(config *app.Config, logger port.Logger) (*app.App, error) {
	if config.Debug {
		err := logger.EnableDebug()
		if err != nil {
			return nil, err
		}
	}

	storage, err := adapter.NewStorage(config.DatabaseURI)
	if err != nil {
		return nil, err
	}

	err = storage.Initialization(config.MigrationDir)
	if err != nil {
		return nil, err
	}

	var (
		router             = adapter.NewRouter()
		cooker             = adapter.NewCooker()
		hasher             = adapter.NewHasher()
		encryptor          = adapter.NewEncryptor(config.CipherPass)
		transactor         = adapter.NewTransactor(storage.Connect())
		userRepository     = repository.NewUserRepository(transactor, storage.Connect())
		orderRepository    = repository.NewOrderRepository(transactor, storage.Connect())
		withdrawRepository = repository.NewWithdrawRepository(transactor, storage.Connect())
	)

	return app.NewApp(
		config,
		logger,
		router,
		encryptor,
		hasher,
		cooker,
		storage,
		transactor,
		userRepository,
		orderRepository,
		withdrawRepository,
	), nil
}
