package usecase

import (
	"net/http"

	"github.com/alrund/yp-1-project/internal/domain/port"
	"github.com/google/uuid"
)

func SetCookie(
	userID uuid.UUID,
	sessionCookieName, sessionCookieDuration string,
	encryptor port.Encryptor,
	cooker port.CookieWithDurationAdder,
	w http.ResponseWriter,
) error {
	encryptedUserID, err := encryptor.Encrypt(userID.String())
	if err != nil {
		return err
	}

	err = cooker.AddCookieWithDuration(sessionCookieName, encryptedUserID, sessionCookieDuration, w)
	if err != nil {
		return err
	}

	return nil
}
