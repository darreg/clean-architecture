package handler

import (
	"net/http"

	"github.com/alrund/yp-1-project/internal/application/app"
)

func LoginHandler(a *app.App) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		a.PlainRespond(w, r, http.StatusOK, []byte("LoginHandler"))
	}

	return http.HandlerFunc(fn)
}
