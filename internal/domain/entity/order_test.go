//go:build unit

package entity

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestHash(t *testing.T) {
	uploadedAt, _ := time.Parse("2006-01-02T15:04:05", "2022-08-05T18:47:28")
	processedAt, _ := time.Parse("2006-01-02T15:04:05", "2022-08-05T19:47:28")

	tests := []struct {
		name string
		data Order
		want string
	}{
		{
			name: "success",
			data: Order{
				Number:      OrderNumber("3272700463"),
				UserID:      uuid.New(),
				Status:      New,
				Accrual:     111.1,
				UploadedAt:  &uploadedAt,
				ProcessedAt: &processedAt,
			},
			want: `{
				"accrual":111.1, 
				"number":"3272700463", 
				"processed_at":"2022-08-05T19:47:28Z", 
				"status":"NEW", 
				"uploaded_at":"2022-08-05T18:47:28Z"
			}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			orderJSON, err := json.Marshal(tt.data)

			assert.NoError(t, err)
			assert.JSONEq(t, tt.want, string(orderJSON))
		})
	}
}
