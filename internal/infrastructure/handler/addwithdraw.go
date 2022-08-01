package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/alrund/yp-1-project/internal/application/app"
	"github.com/alrund/yp-1-project/internal/application/usecase"
	"github.com/alrund/yp-1-project/internal/domain/entity"
	"github.com/alrund/yp-1-project/internal/infrastructure/helper"
)

type WithdrawRequest struct {
	Order string
	Sum   float32
}

func AddWithdrawHandler(a *app.App) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		userID, err := a.AuthRequired(r)
		if err != nil {
			a.Warn(w, r, http.StatusUnauthorized, usecase.ErrNotAuthenticated)

			return
		}

		if !helper.HasContentType(r, "application/json") {
			a.Warn(w, r, http.StatusBadRequest, usecase.ErrInvalidRequestFormat)

			return
		}

		var withdrawRequest WithdrawRequest
		if err := json.NewDecoder(r.Body).Decode(&withdrawRequest); err != nil {
			a.Warn(w, r, http.StatusBadRequest, usecase.ErrInvalidRequestFormat)

			return
		}

		err = usecase.AddWithdraw(
			r.Context(),
			withdrawRequest.Order,
			withdrawRequest.Sum,
			userID,
			a.UserRepository,
			a.WithdrawRepository,
			a.Transactor,
		)
		if err != nil {
			switch {
			case errors.Is(err, entity.ErrInvalidOrderFormat):
				a.Warn(w, r, http.StatusUnprocessableEntity, entity.ErrInvalidOrderFormat)
			case errors.Is(err, usecase.ErrNotEnoughFunds):
				a.Warn(w, r, http.StatusPaymentRequired, usecase.ErrNotEnoughFunds)
			default:
				a.Error(w, r, http.StatusInternalServerError, usecase.ErrInternalServerError)
			}

			return
		}

		a.PlainRespond(w, r, http.StatusOK, []byte(http.StatusText(http.StatusOK)))
	}

	return http.HandlerFunc(fn)
}
