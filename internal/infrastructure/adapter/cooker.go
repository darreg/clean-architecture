package adapter

import (
	"net/http"
	"time"
)

type Cooker struct{}

func NewCooker() *Cooker {
	return &Cooker{}
}

func (c Cooker) ClearCookie(name string, w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:     name,
		Value:    "",
		Path:     "/",
		Expires:  time.Unix(0, 0),
		MaxAge:   0,
		HttpOnly: true,
	})
}

func (c Cooker) AddCookie(name, value string, expireTime time.Time, w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:     name,
		Value:    value,
		Path:     "/",
		Expires:  expireTime,
		HttpOnly: true,
	})
}

func (c Cooker) AddCookieWithDuration(name, value, duration string, w http.ResponseWriter) error {
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
