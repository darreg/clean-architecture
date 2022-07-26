package middleware

import (
	"context"
	"errors"
	"net/http"


	"github.com/alrund/yp-1-project/internal/application/app"
	"github.com/alrund/yp-1-project/internal/application/usecase"
	"github.com/alrund/yp-1-project/internal/domain/port"
)

type ContextKey string

func Auth(a *app.App) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			sessionCookieName := ContextKey(a.Config.SessionCookieName)

			userID, err := getCookie(r, a.Encryption, sessionCookieName)
			if err != nil && !errors.Is(err, http.ErrNoCookie) {
				a.Error(w, r, http.StatusInternalServerError, usecase.ErrInternalServerError)
			}

			ctx := context.WithValue(r.Context(), sessionCookieName, userID)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func getCookie(r *http.Request, enc port.Encryptor, sessionCookieName ContextKey) (string, error) {
	sessionCookie, err := r.Cookie(string(sessionCookieName))
	if err != nil {
		return "", err
	}

	if sessionCookie.Value == "" {
		return "", nil
	}

	userID, err := enc.Decrypt(sessionCookie.Value)
	if err != nil {
		return "", err
	}

	return userID, nil
}
