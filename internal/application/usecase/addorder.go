package usecase

import (
	"errors"
	"strconv"
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

	if !isValidLuhn(number) {
		return ErrInvalidOrderFormat
	}

	uploadedAt := time.Now()
	err = orderRepository.Add(&entity.Order{
		Number:     number,
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

func isValidLuhn(number string) bool {
	numberRunes := []rune(number)

	sum, err := strconv.Atoi(string(numberRunes[len(number)-1]))
	if err != nil {
		return false
	}
	parity := len(number) % 2
	for i := len(number) - 2; i >= 0; i-- {
		summand, err := strconv.Atoi(string(numberRunes[i]))
		if err != nil {
			return false
		}

		if i%2 == parity {
			product := summand * 2
			if product > 9 {
				summand = product - 9
			} else {
				summand = product
			}
		}
		sum += summand
	}
	return (sum % 10) == 0
}
