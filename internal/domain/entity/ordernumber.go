package entity

import (
	"errors"
	"strconv"
)

var ErrInvalidOrderFormat = errors.New("invalid order number format")

type OrderNumber string

func (o OrderNumber) String() string {
	return string(o)
}

func NewOrderNumber(number string) (*OrderNumber, error) {
	if !isValidLuhn(number) {
		return nil, ErrInvalidOrderFormat
	}

	n := OrderNumber(number)
	return &n, nil
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
