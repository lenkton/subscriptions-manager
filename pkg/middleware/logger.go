package middleware

import (
	"log"
	"net/http"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("INFO: processing %v %v", r.Method, r.URL)
		spy := &spyWriter{ResponseWriter: w}
		next.ServeHTTP(spy, r)
		log.Printf("INFO: responded to %v %v with %v", r.Method, r.URL, spy.code)
	})
}

type spyWriter struct {
	http.ResponseWriter
	code int
}

func (spy *spyWriter) WriteHeader(code int) {
	spy.code = code
	spy.ResponseWriter.WriteHeader(code)
}
