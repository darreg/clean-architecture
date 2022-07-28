package app

import "flag"

type Flags struct {
	Debug bool
	A     string
	D     string
	R     string
}

func NewFlags() *Flags {
	flags := &Flags{}

	flag.BoolVar(&flags.Debug, "debug", false, "Enable debug")
	flag.StringVar(&flags.A, "a", "", "Run address")
	flag.StringVar(&flags.D, "d", "", "Database uri")
	flag.StringVar(&flags.R, "r", "", "Accrual system address")

	flag.Parse()

	return flags
}
