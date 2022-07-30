package handler

import (
	"context"
	"errors"
	"io"
	"net/http"
	"time"

	"github.com/alrund/yp-1-project/internal/application/app"
	"github.com/alrund/yp-1-project/internal/application/usecase"
	"github.com/alrund/yp-1-project/internal/domain/entity"
	"github.com/alrund/yp-1-project/internal/infrastructure/helper"
)

func AddOrderHandler(a *app.App) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		userID, err := a.AuthRequired(r)
		if err != nil {
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

		number := string(b)

		err = usecase.AddOrder(
			r.Context(),
			number,
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

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
		defer cancel()

		go usecase.Accrual(
			ctx,
			number, userID, a.Config.AccrualSystemAddress, a.Config.AccrualSystemMethod,
			a.UserRepository,
			a.OrderRepository,
			a.Logger,
		)

		a.PlainRespond(w, r, http.StatusAccepted, []byte(http.StatusText(http.StatusAccepted)))
	}

	return http.HandlerFunc(fn)
}
