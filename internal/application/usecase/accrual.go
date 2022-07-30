package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/alrund/yp-1-project/internal/domain/entity"
	"github.com/alrund/yp-1-project/internal/domain/port"
	"github.com/google/uuid"
)

func Accrual(
	ctx context.Context,
	number, userID, accrualSystemAddress, accrualSystemMethod string,
	userRepository port.UserRepository,
	orderRepository port.OrderRepository,
	logger port.Logger,
) {
	userUUID, err := uuid.Parse(userID)
	if err != nil {
		logger.Error(err)
		return
	}

	user, err := userRepository.Get(ctx, userUUID)
	if err != nil {
		logger.Error(err)
		return
	}

	for {
		select {
		case <-ctx.Done():
			logger.Error(fmt.Errorf("accrual canceled by context"))
			return
		default:
			logger.Info("Accrual request", "number", number)

			accrualResult, err := AccrualRequest(ctx, number, accrualSystemAddress, accrualSystemMethod)
			if err != nil {
				logger.Error(err)
				return
			}

			logger.Info("Accrual response", "number", accrualResult.OrderNumber, "status", accrualResult.Status.String())

			err = AccrualProcess(ctx, accrualResult, user, userRepository, orderRepository)
			if err != nil {
				logger.Error(err)
				return
			}

			if accrualResult.Status == entity.Invalid || accrualResult.Status == entity.Processed {
				return
			}

			time.Sleep(time.Second)
		}
	}
}
