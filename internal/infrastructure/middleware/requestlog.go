package middleware

import (
	"net/http"

	"github.com/alrund/yp-1-project/internal/application/app"
)

func RequestLog(a *app.App) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			a.Logger.Info("request",
				"method", r.Method,
				"path", r.URL.Path,
				"remoteAddr", r.RemoteAddr,
				"userAgent", r.UserAgent(),
			)
			next.ServeHTTP(w, r)
		})
	}
}
