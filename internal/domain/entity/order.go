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
	var ProcessedAt string
	if o.ProcessedAt != nil {
		ProcessedAt = o.ProcessedAt.Format(time.RFC3339)
	}

	type OrderView struct {
		Number      OrderNumber `json:"number"`
		Status      string      `json:"status"`
		Accrual     float32     `json:"accrual,omitempty"`
		UploadedAt  string      `json:"uploaded_at"`
		ProcessedAt string      `json:"processed_at,omitempty"`
	}

	return json.Marshal(&OrderView{
		Number:      o.Number,
		Status:      o.Status.String(),
		Accrual:     o.Accrual,
		UploadedAt:  o.UploadedAt.Format(time.RFC3339),
		ProcessedAt: ProcessedAt,
	})
}
