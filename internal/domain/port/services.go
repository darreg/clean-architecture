package port

import (
	"net/http"
)

type Logger interface {
	Debug(msg string)
	Warn(msg string)
	Info(msg string)
	Error(err error)
	Fatal(err error)
}

type Router interface {
	ServeHTTP(wrt http.ResponseWriter, req *http.Request)
	WithPrefix(prefix string) Router
	Get(path string, handler http.Handler)
	Post(path string, handler http.Handler)
	Use(mwf ...func(http.Handler) http.Handler)
}
