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
	switch s {
	case New:
		return "NEW"
	case Processing:
		return "PROCESSING"
	case Invalid:
		return "INVALID"
	case Processed:
		return "PROCESSED"
	default:
		return ""
	}
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
