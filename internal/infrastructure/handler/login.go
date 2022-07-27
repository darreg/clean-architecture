package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/alrund/yp-1-project/internal/application/app"
	"github.com/alrund/yp-1-project/internal/application/usecase"
)

func LoginHandler(a *app.App) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		var cred usecase.Credential
		if err := json.NewDecoder(r.Body).Decode(&cred); err != nil {
			a.Error(w, r, http.StatusBadRequest, usecase.ErrInvalidRequestFormat)

			return
		}

		err := usecase.Login(
			cred,
			a.Config.SessionCookieName,
			a.Config.SessionCookieDuration,
			a.UserRepository,
			a.Encryptor,
			a.Cooker,
			w,
		)
		if err != nil {
			switch {
			case errors.Is(err, usecase.ErrUserNotFound):
				a.Error(w, r, http.StatusUnauthorized, usecase.ErrNotAuthenticated)
			default:
				a.Error(w, r, http.StatusInternalServerError, usecase.ErrInternalServerError)
			}
		}

		a.JSONRespond(w, r, http.StatusOK, "Ok")
	}

	return http.HandlerFunc(fn)
}
