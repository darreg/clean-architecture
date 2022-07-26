package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/alrund/yp-1-project/internal/application/app"
	"github.com/alrund/yp-1-project/internal/application/usecase"
)

func RegisterHandler(a *app.App) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		var regData usecase.RegistrationData
		if err := json.NewDecoder(r.Body).Decode(&regData); err != nil {
			a.Error(w, r, http.StatusBadRequest, usecase.ErrInvalidRequestFormat)

			return
		}

		err := usecase.Registration(
			regData,
			a.Config.SessionCookieName,
			a.Config.SessionCookieDuration,
			a.UserRepository,
			a.Encryption,
			a.Cooker,
			w,
		)
		if err != nil {
			switch {
			case errors.Is(err, usecase.ErrLoginAlreadyUse):
				a.Error(w, r, http.StatusConflict, usecase.ErrLoginAlreadyUse)
			default:
				a.Error(w, r, http.StatusInternalServerError, usecase.ErrInternalServerError)
			}
		}

		a.JSONRespond(w, r, http.StatusOK, "Ok")
	}

	return http.HandlerFunc(fn)
}
