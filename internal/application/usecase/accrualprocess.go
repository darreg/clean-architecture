package usecase

import (
	"context"

	"github.com/alrund/yp-1-project/internal/domain/entity"
	"github.com/alrund/yp-1-project/internal/domain/port"
)

func AccrualProcess(
	ctx context.Context,
	accrualResult *AccrualResult,
	user *entity.User,
	userRepository port.UserRepository,
	orderRepository port.OrderRepository,
) error {
	err := orderRepository.WithinTransaction(ctx, func(txCtx context.Context) error {
		order, err := orderRepository.Get(ctx, accrualResult.OrderNumber)
		if err != nil {
			return err
		}

		order.Status = accrualResult.Status
		if accrualResult.Status == entity.Processed && order.Accrual > 0 {
			order.Accrual += accrualResult.Accrual
			err = userRepository.Accrual(txCtx, user, accrualResult.Accrual)
			if err != nil {
				return err
			}
		}

		err = orderRepository.Change(txCtx, order)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}
	return nil
}