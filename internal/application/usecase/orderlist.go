package usecase

import (
	"context"

	"github.com/alrund/yp-1-project/internal/domain/entity"
	"github.com/alrund/yp-1-project/internal/domain/port"
	"github.com/google/uuid"
)

func OrderList(
	ctx context.Context,
	userID string,
	orderRepository port.OrderAllByUserGetter,
	userRepository port.UserGetter,
) ([]*entity.Order, error) {
	userUUID, err := uuid.Parse(userID)
	if err != nil {
		return nil, err
	}

	user, err := userRepository.Get(ctx, userUUID)
	if err != nil {
		return nil, err
	}

	orders, err := orderRepository.GetAllByUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return orders, nil
}
