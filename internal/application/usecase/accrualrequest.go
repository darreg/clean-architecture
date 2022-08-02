package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/alrund/yp-1-project/internal/domain/port"
	"net/http"
	"time"

	"github.com/alrund/yp-1-project/internal/domain/entity"
)

const (
	ConnectMaxWaitTime = 1 * time.Second
	RequestMaxWaitTime = 5 * time.Second
)

type AccrualResult struct {
	OrderNumber string             `json:"order"`
	Status      entity.OrderStatus `json:"status"`
	Accrual     float32            `json:"accrual,omitempty"`
}

func (a *AccrualResult) UnmarshalJSON(data []byte) error {
	type Alias AccrualResult
	res := &struct {
		*Alias
		Status string `json:"status"`
	}{
		Alias: (*Alias)(a),
	}
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	status, err := entity.ToOrderStatus(res.Status)
	if err != nil {
		return err
	}

	a.Status = status

	return nil
}

func AccrualRequest(
	ctx context.Context,
	number, accrualSystemAddress, accrualSystemMethod string,
	logger port.Logger,
) (*AccrualResult, error) {
	var result AccrualResult

	ctx, cancel := context.WithTimeout(ctx, RequestMaxWaitTime)
	defer cancel()

	logger.Info("Accrual request", "number", number)

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		accrualSystemAddress+fmt.Sprintf(accrualSystemMethod, number),
		nil,
	)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	client.Timeout = ConnectMaxWaitTime

	rsp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer rsp.Body.Close()

	err = json.NewDecoder(rsp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	logger.Info("Accrual response", "number", result.OrderNumber, "status", result.Status.String())

	return &result, nil
}
