package usecase

import (
	"github.com/alrund/yp-1-project/internal/domain/entity"
	"github.com/alrund/yp-1-project/internal/domain/port"
	"github.com/google/uuid"
)

func WithdrawList(
	userID string,
	withdrawRepository port.WithdrawRepository,
	userRepository port.UserRepository,
) ([]*entity.Withdraw, error) {
	userUUID, err := uuid.Parse(userID)
	if err != nil {
		return nil, err
	}

	user, err := userRepository.Get(userUUID)
	if err != nil {
		return nil, err
	}

	withdraws, err := withdrawRepository.GetAllByUser(user)
	if err != nil {
		return nil, err
	}

	return withdraws, nil
}
