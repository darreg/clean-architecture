package middleware

import (
	"net/http"

	"github.com/alrund/yp-1-project/internal/application/app"
)

func RequestLog(a *app.App) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			a.Logger.Info("Request",
				"Method", r.Method,
				"Path", r.URL.Path,
				"RemoteAddr", r.RemoteAddr,
				"UserAgent", r.UserAgent(),
			)
			next.ServeHTTP(w, r)
		})
	}
}
