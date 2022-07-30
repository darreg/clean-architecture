package app

import (
	"github.com/alrund/yp-1-project/internal/domain/port"
)

type Config struct {
	Debug                 bool   `env-default:"false"`
	RunAddress            string `env:"RUN_ADDRESS" env-default:"localhost:3000"`
	DatabaseURI           string `env:"DATABASE_URI" env-default:"postgres://dev:dev@localhost:5432/dev?sslmode=disable"`
	AccrualSystemAddress  string `env:"ACCRUAL_SYSTEM_ADDRESS" env-default:"http://localhost:8080"`
	AccrualSystemMethod   string `env-default:"/api/orders/%s"`
	CipherPass            string `env:"CIPHER_PASSWORD" env-default:"J53RPX6"`
	SessionCookieDuration string `env:"COOKIE_DURATION" env-default:"24h"`
	SessionCookieName     string `env:"COOKIE_NAME" env-default:"sessionID"`
}

func NewConfig(loader port.ConfigLoader) (*Config, error) {
	cfg := &Config{}

	flags := NewFlags()
	cfg.Debug = flags.Debug
	cfg.RunAddress = flags.A
	cfg.DatabaseURI = flags.D
	cfg.AccrualSystemAddress = flags.R

	if err := loader.Load(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
