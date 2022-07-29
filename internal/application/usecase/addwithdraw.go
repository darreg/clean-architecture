package usecase

import (
	"time"

	"github.com/alrund/yp-1-project/internal/domain/entity"
	"github.com/alrund/yp-1-project/internal/domain/port"
	"github.com/google/uuid"
)

func AddWithdraw(
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

	user, err := userRepository.Get(userUUID)
	if err != nil {
		return err
	}

	order, err := orderRepository.Get(number)
	if err != nil {
		return err
	}

	if user.Current < sum {
		return ErrNotEnoughFunds
	}

	processedAt := time.Now()
	err = withdrawRepository.Add(&entity.Withdraw{
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
