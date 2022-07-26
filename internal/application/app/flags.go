package app

import "flag"

type Flags struct {
	A string
	D string
	R string
}

func NewFlags() *Flags {
	flags := &Flags{}

	flag.StringVar(&flags.A, "a", "", "Run address")
	flag.StringVar(&flags.D, "d", "", "Database uri")
	flag.StringVar(&flags.R, "r", "", "Accrual system address")

	flag.Parse()

	return flags
}
