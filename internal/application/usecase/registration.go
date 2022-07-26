package usecase

import (
	"errors"
	"net/http"


	"github.com/alrund/yp-1-project/internal/domain/port"
	"github.com/google/uuid"
)

type RegistrationData struct {
	Login    string
	Password string
}

func Registration(
	regData RegistrationData,
	sessionCookieName, sessionCookieDuration string,
	repository port.UserRepository,
	encryptor port.Encryptor,
	cooker port.Cooker,
	w http.ResponseWriter,
) error {
	user, err := repository.FindByLogin(regData.Login)
	if err != nil && !errors.Is(err, ErrUserNotFound) {
		return err
	}

	if user != nil {
		return ErrLoginAlreadyUse
	}

	user, err = repository.Create(uuid.New(), regData.Login, regData.Password)
	if err != nil {
		return err
	}

	encryptedUserID, err := encryptor.Encrypt(user.ID.String())
	if err != nil {
		return err
	}

	err = cooker.AddCookieWithDuration(sessionCookieName, encryptedUserID, sessionCookieDuration, w)
	if err != nil {
		return err
	}

	return nil
}