package entity

const (
	New OrderStatus = iota
	Processing
	Invalid
	Processed
)

type OrderStatus int

func (s OrderStatus) String() string {
	statuses := [...]string{"NEW", "PROCESSING", "INVALID", "PROCESSED"}
	if len(statuses) < int(s) {
		return ""
	}
	return statuses[s]
}
