package service

import (
	"github.com/alrund/yp-1-project/internal/application/app"
	"github.com/alrund/yp-1-project/internal/domain/port"
	"github.com/ilyakaznacheev/cleanenv"
)

func GetConfig(flags *app.Flags, logger port.Logger) *app.Config {
	cfg := &app.Config{}

	cfg.RunAddress = flags.A
	cfg.DatabaseURI = flags.D
	cfg.AccrualSystemAddress = flags.R

	if err := cleanenv.ReadEnv(cfg); err != nil {
		logger.Fatal(err)
	}

	return cfg
}
