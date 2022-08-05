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
	number, userID, accrualSystemAddress, accrualSystemMethod, accrualSystemPollInterval string,
	userRepository port.UserRepository,
	orderRepository port.OrderRepository,
	transactor port.Transactor,
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

	pollInterval, err := time.ParseDuration(accrualSystemPollInterval)
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
			accrualResult, err := AccrualRequest(ctx, number, accrualSystemAddress, accrualSystemMethod, logger)
			if err != nil {
				logger.Error(err)
				return
			}

			err = AccrualProcess(ctx, accrualResult, user, userRepository, orderRepository, transactor)
			if err != nil {
				logger.Error(err)
				return
			}

			if accrualResult.Status == entity.Invalid || accrualResult.Status == entity.Processed {
				return
			}

			time.Sleep(pollInterval)
		}
	}
}
