package middleware

import "net/http"

// TODO: write my own http.Error(): the default one breaks the content type header
func JSONEndpoint(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")
		next.ServeHTTP(w, r)
	})
}
