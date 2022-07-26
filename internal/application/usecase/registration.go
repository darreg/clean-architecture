package usecase

import (
	"net/http"

	"github.com/alrund/yp-1-project/internal/domain/port"
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
	if err != nil {
		return err
	}

	if user != nil {
		return ErrLoginAlreadyUse
	}

	user, err = repository.Create(regData.Login, regData.Password)
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
