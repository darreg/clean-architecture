package handler

import (
	"errors"
	"io"
	"net/http"

	"github.com/alrund/yp-1-project/internal/application/app"
	"github.com/alrund/yp-1-project/internal/application/usecase"
	"github.com/alrund/yp-1-project/internal/domain/entity"
	"github.com/alrund/yp-1-project/internal/infrastructure/helper"
	"github.com/alrund/yp-1-project/internal/infrastructure/middleware"
)

func AddOrderHandler(a *app.App) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		contextUserID := r.Context().Value(middleware.ContextKey(a.Config.SessionCookieName))
		userID, ok := contextUserID.(string)
		if !ok || userID == "" {
			a.Warn(w, r, http.StatusUnauthorized, usecase.ErrNotAuthenticated)

			return
		}

		b, err := io.ReadAll(r.Body)
		if err != nil {
			a.Error(w, r, http.StatusInternalServerError, usecase.ErrInternalServerError)

			return
		}

		if len(b) == 0 || !helper.HasContentType(r, "text/plain") {
			a.Warn(w, r, http.StatusBadRequest, usecase.ErrInvalidRequestFormat)

			return
		}

		err = usecase.AddOrder(
			string(b),
			userID,
			a.OrderRepository,
			a.UserRepository,
		)
		if err != nil {
			switch {
			case errors.Is(err, usecase.ErrOrderAlreadyUploaded):
				a.Warn(w, r, http.StatusOK, usecase.ErrOrderAlreadyUploaded)
			case errors.Is(err, usecase.ErrOrderAlreadyUploadedAnotherUser):
				a.Warn(w, r, http.StatusConflict, usecase.ErrOrderAlreadyUploadedAnotherUser)
			case errors.Is(err, entity.ErrInvalidOrderFormat):
				a.Warn(w, r, http.StatusUnprocessableEntity, entity.ErrInvalidOrderFormat)
			default:
				a.Error(w, r, http.StatusInternalServerError, usecase.ErrInternalServerError)
			}

			return
		}

		a.PlainRespond(w, r, http.StatusAccepted, []byte(http.StatusText(http.StatusAccepted)))
	}

	return http.HandlerFunc(fn)
}
