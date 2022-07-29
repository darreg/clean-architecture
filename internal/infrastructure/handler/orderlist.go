package handler

import (
	"errors"
	"net/http"

	"github.com/alrund/yp-1-project/internal/application/app"
	"github.com/alrund/yp-1-project/internal/application/usecase"
)

func OrderListHandler(a *app.App) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		userID, err := a.AuthRequired(r)
		if err != nil {
			a.Warn(w, r, http.StatusUnauthorized, usecase.ErrNotAuthenticated)

			return
		}

		orders, err := usecase.OrderList(
			r.Context(),
			userID,
			a.OrderRepository,
			a.UserRepository,
		)
		if err != nil {
			switch {
			case errors.Is(err, usecase.ErrOrderNotFound):
				a.Warn(w, r, http.StatusNoContent, usecase.ErrOrderNotFound)
			default:
				a.Error(w, r, http.StatusInternalServerError, usecase.ErrInternalServerError)
			}

			return
		}

		a.JSONRespond(w, r, http.StatusOK, orders)
	}

	return http.HandlerFunc(fn)
}
