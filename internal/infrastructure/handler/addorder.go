package handler

import (
	"errors"
	"io"
	"mime"
	"net/http"

	"github.com/alrund/yp-1-project/internal/application/app"
	"github.com/alrund/yp-1-project/internal/application/usecase"
	"github.com/alrund/yp-1-project/internal/infrastructure/middleware"
)

func AddOrderHandler(a *app.App) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		contextUserID := r.Context().Value(middleware.ContextKey(a.Config.SessionCookieName))
		userID, ok := contextUserID.(string)
		if !ok || userID == "" {
			a.Error(w, r, http.StatusUnauthorized, usecase.ErrNotAuthenticated)

			return
		}

		b, err := io.ReadAll(r.Body)
		if err != nil {
			a.Error(w, r, http.StatusInternalServerError, usecase.ErrInternalServerError)

			return
		}

		if len(b) == 0 || !hasContentType(r, "text/plain") {
			a.Error(w, r, http.StatusBadRequest, usecase.ErrInvalidRequestFormat)

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
				a.Error(w, r, http.StatusOK, usecase.ErrOrderAlreadyUploaded)
			case errors.Is(err, usecase.ErrOrderAlreadyUploadedAnotherUser):
				a.Error(w, r, http.StatusConflict, usecase.ErrOrderAlreadyUploadedAnotherUser)
			case errors.Is(err, usecase.ErrInvalidOrderFormat):
				a.Error(w, r, http.StatusUnprocessableEntity, usecase.ErrInvalidOrderFormat)
			default:
				a.Error(w, r, http.StatusInternalServerError, usecase.ErrInternalServerError)
			}

			return
		}

		a.JSONRespond(w, r, http.StatusAccepted, "Accepted")
	}

	return http.HandlerFunc(fn)
}

func hasContentType(r *http.Request, mimetype string) bool {
	contentType := r.Header.Get("Content-type")
	t, _, err := mime.ParseMediaType(contentType)
	if err == nil && t == mimetype {
		return true
	}
	return false
}
