package adapter

import (
	"net/http"

	"github.com/alrund/yp-1-project/internal/domain/port"
	"github.com/gorilla/mux"
)

type Router struct {
	mux *mux.Router
}

func NewRouter() *Router {
	return &Router{
		mux: mux.NewRouter(),
	}
}

func (r *Router) ServeHTTP(wrt http.ResponseWriter, req *http.Request) {
	r.mux.ServeHTTP(wrt, req)
}

func (r *Router) WithPrefix(prefix string) port.Router {
	return &Router{
		mux: r.mux.PathPrefix(prefix).Subrouter(),
	}
}

func (r *Router) Get(path string, handler http.Handler) {
	r.mux.Handle(path, handler).Methods(http.MethodGet)
}

func (r *Router) Post(path string, handler http.Handler) {
	r.mux.Handle(path, handler).Methods(http.MethodPost)
}

func (r *Router) Use(mwf ...func(http.Handler) http.Handler) {
	for _, fn := range mwf {
		r.mux.Use(fn)
	}
}
