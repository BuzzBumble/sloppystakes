package main

import (
	"crypto/rand"
	"net/http"

	"github.com/gorilla/csrf"
)

func WriteToConsole(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}

func GenerateCSRF(next http.Handler) http.Handler {
	key := make([]byte, 64)

	_, _ = rand.Read(key)
	CSRF := csrf.Protect(key, csrf.Secure(app.IsSecure))
	return CSRF(next)
}

func SessionLoad(next http.Handler) http.Handler {
	return sm.LoadAndSave(next)
}
