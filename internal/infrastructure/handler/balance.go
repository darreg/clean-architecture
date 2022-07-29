package handler

import (
	"net/http"

	"github.com/alrund/yp-1-project/internal/application/app"
	"github.com/alrund/yp-1-project/internal/application/usecase"
)

func BalanceHandler(a *app.App) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		userID, err := a.AuthRequired(r)
		if err != nil {
			a.Warn(w, r, http.StatusUnauthorized, usecase.ErrNotAuthenticated)

			return
		}
		balance, err := usecase.CheckBalance(
			r.Context(),
			userID,
			a.UserRepository,
			a.WithdrawRepository,
		)
		if err != nil {
			a.Error(w, r, http.StatusInternalServerError, usecase.ErrInternalServerError)

			return
		}

		a.JSONRespond(w, r, http.StatusOK, balance)
	}

	return http.HandlerFunc(fn)
}
