package handler

import (
	"errors"
	"net/http"

	"github.com/alrund/yp-1-project/internal/application/app"
	"github.com/alrund/yp-1-project/internal/application/usecase"
)

func WithdrawListHandler(a *app.App) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		userID, err := a.AuthRequired(r)
		if err != nil {
			a.Warn(w, r, http.StatusUnauthorized, usecase.ErrNotAuthenticated)

			return
		}

		orders, err := usecase.WithdrawList(
			r.Context(),
			userID,
			a.WithdrawRepository,
			a.UserRepository,
		)
		if err != nil {
			switch {
			case errors.Is(err, usecase.ErrWithdrawNotFound):
				a.Warn(w, r, http.StatusNoContent, usecase.ErrWithdrawNotFound)
			default:
				a.Error(w, r, http.StatusInternalServerError, usecase.ErrInternalServerError)
			}

			return
		}

		a.JSONRespond(w, r, http.StatusOK, orders)
	}

	return http.HandlerFunc(fn)
}
