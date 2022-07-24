package service

import (
	"flag"

	"github.com/alrund/yp-1-project/internal/application/app"
)

func GetFlags() *app.Flags {
	flags := &app.Flags{}

	flag.StringVar(&flags.A, "a", "", "Run address")
	flag.StringVar(&flags.D, "d", "", "Database uri")
	flag.StringVar(&flags.R, "r", "", "Accrual system address")

	flag.Parse()

	return flags
}
