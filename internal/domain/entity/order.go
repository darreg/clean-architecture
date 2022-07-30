package entity

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type Order struct {
	Number      OrderNumber `json:"number"`
	UserID      uuid.UUID   `json:"-"`
	Status      OrderStatus `json:"status"`
	Accrual     float32     `json:"accrual,omitempty"`
	UploadedAt  *time.Time  `json:"uploaded_at"`
	ProcessedAt *time.Time  `json:"processed_at,omitempty"`
}

func (o Order) MarshalJSON() ([]byte, error) {
	type Alias Order

	var ProcessedAt string
	if o.ProcessedAt != nil {
		ProcessedAt = o.ProcessedAt.Format(time.RFC3339)
	}

	aliasValue := &struct {
		Alias
		Number      string `json:"number"`
		UploadedAt  string `json:"uploaded_at"`
		ProcessedAt string `json:"processed_at,omitempty"`
		OrderStatus string `json:"status"`
	}{
		Alias:       Alias(o),
		Number:      o.Number.String(),
		UploadedAt:  o.UploadedAt.Format(time.RFC3339),
		ProcessedAt: ProcessedAt,
		OrderStatus: o.Status.String(),
	}

	return json.Marshal(aliasValue)
}
