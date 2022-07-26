package app

import "github.com/ilyakaznacheev/cleanenv"

type Config struct {
	RunAddress            string `env:"RUN_ADDRESS" env-default:"localhost:3000"`
	DatabaseURI           string `env:"DATABASE_URI" env-default:"postgres://dev:dev@localhost:5432/dev?sslmode=disable"`
	AccrualSystemAddress  string `env:"ACCRUAL_SYSTEM_ADDRESS" env-default:""`
	CipherPass            string `env:"CIPHER_PASSWORD" env-default:"J53RPX6"`
	SessionCookieDuration string `env:"COOKIE_DURATION" env-default:"24h"`
	SessionCookieName     string `env:"COOKIE_NAME" env-default:"sessionID"`
}

func NewConfig() (*Config, error) {
	cfg := &Config{}

	flags := NewFlags()
	cfg.RunAddress = flags.A
	cfg.DatabaseURI = flags.D
	cfg.AccrualSystemAddress = flags.R

	if err := cleanenv.ReadEnv(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
