package entity

import "errors"

const (
	New OrderStatus = iota
	Processing
	Invalid
	Processed
)

var ErrInvalidOrderStatus = errors.New("invalid order status")

type OrderStatus int

func (s OrderStatus) String() string {
	statuses := [...]string{"NEW", "PROCESSING", "INVALID", "PROCESSED"}
	if len(statuses) < int(s) {
		return ""
	}
	return statuses[s]
}

func ToOrderStatus(str string) (OrderStatus, error) {
	switch str {
	case "NEW", "REGISTERED":
		return New, nil
	case "PROCESSING":
		return Processing, nil
	case "INVALID":
		return Invalid, nil
	case "PROCESSED":
		return Processed, nil
	default:
		return -1, ErrInvalidOrderStatus
	}
}
