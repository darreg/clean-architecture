package app

type Config struct {
	RunAddress           string `env:"RUN_ADDRESS" env-default:"localhost:3000"`
	DatabaseURI          string `env:"DATABASE_URI" env-default:"postgres://dev:dev@localhost:5432/dev"`
	AccrualSystemAddress string `env:"ACCRUAL_SYSTEM_ADDRESS" env-default:""`
}
