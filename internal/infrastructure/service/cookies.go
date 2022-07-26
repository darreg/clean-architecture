package service

import (
	"net/http"
	"time"
)

type Cookies struct{}

func NewCookies() *Cookies {
	return &Cookies{}
}

func (c Cookies) ClearCookie(name string, w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:     name,
		Value:    "",
		Path:     "/",
		Expires:  time.Unix(0, 0),
		MaxAge:   0,
		HttpOnly: true,
	})
}

func (c Cookies) AddCookie(name, value string, expireTime time.Time, w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:     name,
		Value:    value,
		Path:     "/",
		Expires:  expireTime,
		HttpOnly: true,
	})
}

func (c Cookies) AddCookieWithDuration(name, value, duration string, w http.ResponseWriter) error {
	expireDuration, err := time.ParseDuration(duration)
	if err != nil {
		return err
	}

	http.SetCookie(w, &http.Cookie{
		Name:     name,
		Value:    value,
		Path:     "/",
		Expires:  time.Now().Add(expireDuration),
		HttpOnly: true,
	})

	return nil
}
