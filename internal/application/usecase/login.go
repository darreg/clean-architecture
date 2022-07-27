package usecase

import (
	"net/http"

	"github.com/alrund/yp-1-project/internal/domain/port"
)

type Credential struct {
	Login    string
	Password string
}

func Login(
	cred Credential,
	sessionCookieName, sessionCookieDuration string,
	repository port.UserRepository,
	encryptor port.Encryptor,
	cooker port.Cooker,
	hasher port.PasswordHasher,
	w http.ResponseWriter,
) error {
	user, err := repository.GetByCredential(cred.Login, hasher.Hash(cred.Password))
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
