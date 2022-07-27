package usecase

import (
	"github.com/alrund/yp-1-project/internal/domain/entity"
	"github.com/alrund/yp-1-project/internal/domain/port"
	"github.com/google/uuid"
)

func OrderList(
	userID string,
	orderRepository port.OrderRepository,
	userRepository port.UserRepository,
) ([]*entity.Order, error) {
	userUUID, err := uuid.Parse(userID)
	if err != nil {
		return nil, err
	}

	user, err := userRepository.Get(userUUID)
	if err != nil {
		return nil, err
	}

	orders, err := orderRepository.GetAllByUser(user)
	if err != nil {
		return nil, err
	}

	return orders, nil
}
