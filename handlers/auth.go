package handlers

import (
	"net/http"
)

var BasicAuthHandler = func(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if !ok || username != "username" || password != "password" {
			{
				w.WriteHeader(403)
				return
			}
		}
		next.ServeHTTP(w, r)
	})
}
