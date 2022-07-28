package usecase

import (
	"errors"
)

var (
	ErrInvalidRequestFormat            = errors.New("invalid request format")
	ErrNotAuthenticated                = errors.New("not authenticated")
	ErrInternalServerError             = errors.New("internal server error")
	ErrUserNotFound                    = errors.New("user not found")
	ErrOrderNotFound                   = errors.New("order not found")
	ErrOrderAlreadyUploadedAnotherUser = errors.New("the order number has already been uploaded by another user")
	ErrOrderAlreadyUploaded            = errors.New("the order number has already been uploaded by this user")
	ErrLoginAlreadyUse                 = errors.New("the login is already in use")
)
