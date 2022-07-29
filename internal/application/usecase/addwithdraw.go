package usecase

import (
	"context"
	"time"

	"github.com/alrund/yp-1-project/internal/domain/entity"
	"github.com/alrund/yp-1-project/internal/domain/port"
	"github.com/google/uuid"
)

func AddWithdraw(
	ctx context.Context,
	number string,
	sum int,
	userID string,
	orderRepository port.OrderRepository,
	userRepository port.UserRepository,
	withdrawRepository port.WithdrawRepository,
) error {
	userUUID, err := uuid.Parse(userID)
	if err != nil {
		return err
	}

	user, err := userRepository.Get(ctx, userUUID)
	if err != nil {
		return err
	}

	order, err := orderRepository.Get(ctx, number)
	if err != nil {
		return err
	}

	if user.Current < sum {
		return ErrNotEnoughFunds
	}

	processedAt := time.Now()
	err = withdrawRepository.Add(ctx, &entity.Withdraw{
		ID:          uuid.New(),
		OrderNumber: order.Number,
		UserID:      user.ID,
		Sum:         sum,
		ProcessedAt: &processedAt,
	})
	if err != nil {
		return err
	}

	return nil
}
