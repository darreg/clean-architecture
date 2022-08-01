package port

import (
	"context"
	"database/sql"
	"net/http"
	"time"
)

type ConfigLoader interface {
	Load(cfg interface{}) error
	LoadFile(path string, cfg interface{}) error
}

type Logger interface {
	EnableDebug() error
	Debug(msg string, args ...interface{})
	Warn(msg string, args ...interface{})
	Info(msg string, args ...interface{})
	Error(err error, args ...interface{})
	Fatal(err error, args ...interface{})
}

type Router interface {
	ServeHTTP(wrt http.ResponseWriter, req *http.Request)
	WithPrefix(prefix string) Router
	Get(path string, handler http.Handler)
	Post(path string, handler http.Handler)
	Use(mwf ...func(http.Handler) http.Handler)
}

type PasswordHasher interface {
	Hash(password string) string
}

type Encryptor interface {
	Encrypt(data string) (string, error)
	Decrypt(encrypted string) (string, error)
}

type CookieWithDurationAdder interface {
	AddCookieWithDuration(name, value, duration string, w http.ResponseWriter) error
}

type CookieAdder interface {
	AddCookie(name, value string, expireTime time.Time, w http.ResponseWriter)
}

type CookieCleaner interface {
	ClearCookie(name string, w http.ResponseWriter)
}

type Cooker interface {
	CookieCleaner
	CookieAdder
	CookieWithDurationAdder
}

type TransactSupporter interface {
	WithinTransaction(ctx context.Context, tFunc func(ctx context.Context) error) error
}

type RowQuerier interface {
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
}

type Querier interface {
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
}

type Execer interface {
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
}

type Transactor interface {
	TransactSupporter
	Querier
	RowQuerier
	Execer
}
