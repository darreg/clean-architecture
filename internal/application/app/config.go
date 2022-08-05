package app

import (
	"github.com/alrund/yp-1-project/internal/domain/port"
)

type Config struct {
	Debug                     bool   `env-default:"false"`
	RunAddress                string `env:"RUN_ADDRESS" env-default:"localhost:3000"`
	DatabaseURI               string `env:"DATABASE_URI" env-default:"postgres://dev:dev@localhost:5432/dev?sslmode=disable"` //nolint
	AccrualSystemAddress      string `env:"ACCRUAL_SYSTEM_ADDRESS" env-default:"http://localhost:8080"`
	AccrualSystemMethod       string `env:"ACCRUAL_SYSTEM_METHOD" env-default:"/api/orders/%s"`
	AccrualSystemPollInterval string `env:"ACCRUAL_SYSTEM_POLL_INTERVAL" env-default:"1s"`
	CipherPass                string `env:"CIPHER_PASSWORD" env-default:"J53RPX6"`
	SessionCookieDuration     string `env:"COOKIE_DURATION" env-default:"24h"`
	SessionCookieName         string `env:"COOKIE_NAME" env-default:"sessionID"`
	MigrationDir              string `env-default:"migrations"`
}

func NewConfig(loader port.ConfigLoader) (*Config, error) {
	cfg := &Config{}

	flags := NewFlags()
	cfg.Debug = flags.Debug
	cfg.RunAddress = flags.RunAddress
	cfg.DatabaseURI = flags.DatabaseURI
	cfg.AccrualSystemAddress = flags.AccrualSystemAddress

	if err := loader.Load(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
