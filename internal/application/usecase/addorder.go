package usecase

import (
	"errors"
	"time"

	"github.com/alrund/yp-1-project/internal/domain/entity"
	"github.com/alrund/yp-1-project/internal/domain/port"
	"github.com/google/uuid"
)

func AddOrder(
	number string,
	userID string,
	orderRepository port.OrderRepository,
	userRepository port.UserRepository,
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
	if err != nil && !errors.Is(err, ErrOrderNotFound) {
		return err
	}

	if order != nil {
		if order.UserID == userUUID {
			return ErrOrderAlreadyUploaded
		}

		return ErrOrderAlreadyUploadedAnotherUser
	}

	orderNumber, err := entity.NewOrderNumber(number)
	if err != nil {
		return err
	}

	uploadedAt := time.Now()
	err = orderRepository.Add(&entity.Order{
		Number:     *orderNumber,
		UserID:     user.ID,
		Status:     entity.New,
		Accrual:    0,
		UploadedAt: &uploadedAt,
	})
	if err != nil {
		return err
	}

	return nil
}
