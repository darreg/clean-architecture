package handler

import (
	"net/http"

	"github.com/alrund/yp-1-project/internal/application/app"
)

func RegisterHandler(a *app.App) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		a.PlainRespond(w, r, http.StatusOK, []byte("RegisterHandler"))
	}

	return http.HandlerFunc(fn)
}
