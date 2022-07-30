package app

import (
	"encoding/json"
	"net/http"

	"github.com/alrund/yp-1-project/internal/application/usecase"
)

type SessionContextKey string

func (a *App) Serve() error {
	a.Logger.Info("starting HTTP server", "addr", a.Config.RunAddress)
	return http.ListenAndServe(a.Config.RunAddress, a.Router)
}

func (a *App) AuthRequired(r *http.Request) (string, error) {
	contextUserID := r.Context().Value(SessionContextKey(a.Config.SessionCookieName))
	userID, ok := contextUserID.(string)
	if !ok || userID == "" {
		return "", usecase.ErrNotAuthenticated
	}

	return userID, nil
}

func (a *App) Warn(w http.ResponseWriter, r *http.Request, code int, err error) {
	a.Logger.Warn(err.Error())

	a.JSONRespond(w, r, code, err.Error())
}

func (a *App) Error(w http.ResponseWriter, r *http.Request, code int, err error) {
	a.Logger.Error(err)

	a.JSONRespond(w, r, code, err.Error())
}

func (a *App) PlainRespond(w http.ResponseWriter, r *http.Request, code int, data []byte) {
	w.Header().Add("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(code)

	if _, err := w.Write(data); err != nil {
		a.Logger.Error(err)

		return
	}
}

func (a *App) JSONRespond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)

	if data != nil {
		err := json.NewEncoder(w).Encode(data)
		if err != nil {
			a.Logger.Error(err)

			return
		}
	}
}
