package entity

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type Withdraw struct {
	ID          uuid.UUID   `json:"-"`
	UserID      uuid.UUID   `json:"-"`
	OrderNumber OrderNumber `json:"order"`
	Sum         int         `json:"sum"`
	ProcessedAt *time.Time  `json:"processed_at"`
}

func (w Withdraw) MarshalJSON() ([]byte, error) {
	type Alias Withdraw

	aliasValue := &struct {
		Alias
		ProcessedAt string `json:"processed_at"`
	}{
		Alias:       Alias(w),
		ProcessedAt: w.ProcessedAt.Format(time.RFC3339),
	}

	return json.Marshal(aliasValue)
}
