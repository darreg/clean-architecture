package app

import "flag"

type Flags struct {
	Debug                bool
	RunAddress           string
	DatabaseURI          string
	AccrualSystemAddress string
}

func NewFlags() *Flags {
	flags := &Flags{}

	flag.BoolVar(&flags.Debug, "debug", false, "Enable debug")
	flag.StringVar(&flags.RunAddress, "a", "", "Run address")
	flag.StringVar(&flags.DatabaseURI, "d", "", "Database uri")
	flag.StringVar(&flags.AccrualSystemAddress, "r", "", "Accrual system address")

	flag.Parse()

	return flags
}
