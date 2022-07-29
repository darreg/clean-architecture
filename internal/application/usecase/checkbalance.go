package usecase

import (
	"context"

	"github.com/alrund/yp-1-project/internal/domain/port"
	"github.com/google/uuid"
)

type Balance struct {
	Current   int `json:"current"`
	Withdrawn int `json:"withdrawn"`
}

func CheckBalance(
	ctx context.Context,
	userID string,
	userRepository port.UserRepository,
	withdrawRepository port.WithdrawRepository,
) (*Balance, error) {
	userUUID, err := uuid.Parse(userID)
	if err != nil {
		return nil, err
	}

	user, err := userRepository.Get(ctx, userUUID)
	if err != nil {
		return nil, err
	}

	withdrawn, err := withdrawRepository.GetWithdrawn(ctx, user)
	if err != nil {
		return nil, err
	}

	return &Balance{Current: user.Current, Withdrawn: withdrawn}, nil
}
