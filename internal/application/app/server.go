package app

import (
	"encoding/json"
	"net/http"
)

func (a *App) Serve() error {
	a.Logger.Info("starting HTTP server", "addr", a.Config.RunAddress)
	return http.ListenAndServe(a.Config.RunAddress, a.Router)
}

func (a *App) Warn(w http.ResponseWriter, r *http.Request, code int, err error) {
	a.Logger.Warn(err.Error())

	a.PlainRespond(w, r, code, []byte(err.Error()))
}

func (a *App) Error(w http.ResponseWriter, r *http.Request, code int, err error) {
	a.Logger.Error(err)

	a.PlainRespond(w, r, code, []byte(err.Error()))
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
