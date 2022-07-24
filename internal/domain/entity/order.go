package entity

import (
	"os/user"
	"time"
)

const (
	New OrderStatus = iota
	Processing
	Invalid
	Processed
)

type OrderStatus int

type Order struct {
	Number      string
	User        *user.User
	Status      OrderStatus
	Accrual     int
	UploadedAt  time.Time
	ProcessedAt time.Time
}
