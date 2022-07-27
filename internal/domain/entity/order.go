package entity

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

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

type Order struct {
	Number      string      `json:"number"`
	UserID      uuid.UUID   `json:"-"`
	Status      OrderStatus `json:"status"`
	Accrual     int         `json:"accrual,omitempty"`
	UploadedAt  *time.Time  `json:"uploaded_at" format:"RFC850"`
	ProcessedAt *time.Time  `json:"processed_at,omitempty" format:"RFC3339"`
}

func (o Order) MarshalJSON() ([]byte, error) {
	type Alias Order

	var ProcessedAt string
	if o.ProcessedAt != nil {
		ProcessedAt = o.ProcessedAt.Format(time.RFC3339)
	}

	aliasValue := &struct {
		Alias
		UploadedAt  string `json:"uploaded_at"`
		ProcessedAt string `json:"processed_at,omitempty"`
		OrderStatus string `json:"status"`
	}{
		Alias:       Alias(o),
		UploadedAt:  o.UploadedAt.Format(time.RFC3339),
		ProcessedAt: ProcessedAt,
		OrderStatus: o.Status.String(),
	}

	return json.Marshal(aliasValue)
}
