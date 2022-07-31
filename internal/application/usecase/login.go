package usecase

import (
	"context"
	"net/http"

	"github.com/alrund/yp-1-project/internal/domain/port"
)

type Credential struct {
	Login    string
	Password string
}

func Login(
	ctx context.Context,
	cred Credential,
	sessionCookieName, sessionCookieDuration string,
	userRepository port.UserByCredentialGetter,
	encryptor port.Encryptor,
	cooker port.CookieWithDurationAdder,
	hasher port.PasswordHasher,
	w http.ResponseWriter,
) error {
	user, err := userRepository.GetByCredential(ctx, cred.Login, hasher.Hash(cred.Password))
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
