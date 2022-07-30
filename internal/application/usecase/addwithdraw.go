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
	sum float32,
	userID string,
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

	orderNumber, err := entity.NewOrderNumber(number)
	if err != nil {
		return err
	}

	if user.Current < sum {
		return ErrNotEnoughFunds
	}

	processedAt := time.Now()
	err = withdrawRepository.WithinTransaction(ctx, func(txCtx context.Context) error {
		err := withdrawRepository.Add(txCtx, &entity.Withdraw{
			ID:          uuid.New(),
			OrderNumber: *orderNumber,
			UserID:      user.ID,
			Sum:         sum,
			ProcessedAt: &processedAt,
		})
		if err != nil {
			return err
		}

		err = userRepository.Withdraw(txCtx, user, sum)
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
