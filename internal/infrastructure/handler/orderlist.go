package handler

import (
	"errors"
	"net/http"

	"github.com/alrund/yp-1-project/internal/application/app"
	"github.com/alrund/yp-1-project/internal/application/usecase"
	"github.com/alrund/yp-1-project/internal/infrastructure/middleware"
)

func OrderListHandler(a *app.App) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		contextUserID := r.Context().Value(middleware.ContextKey(a.Config.SessionCookieName))
		userID, ok := contextUserID.(string)
		if !ok {
			a.Error(w, r, http.StatusUnauthorized, usecase.ErrNotAuthenticated)
		}

		orders, err := usecase.OrderList(
			userID,
			a.OrderRepository,
			a.UserRepository,
		)
		if err != nil {
			switch {
			case errors.Is(err, usecase.ErrOrderNotFound):
				a.Error(w, r, http.StatusNoContent, usecase.ErrOrderNotFound)
			default:
				a.Error(w, r, http.StatusInternalServerError, usecase.ErrInternalServerError)
			}

			return
		}

		a.JSONRespond(w, r, http.StatusAccepted, orders)
	}

	return http.HandlerFunc(fn)
}
