package httputil

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func EncodeJSON(w http.ResponseWriter, body any, code int) {
	w.WriteHeader(code)
	err := json.NewEncoder(w).Encode(body)
	if err != nil {
		log.Printf("ERROR: encoding subs: %v\n", err)
	}
}

func DecodeJSON[T any](body io.Reader) (*T, error) {
	var t T
	err := json.NewDecoder(body).Decode(&t)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func WithJSONBody[T any](next http.Handler, key any) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		dto, err := DecodeJSON[T](r.Body)
		if err != nil {
			// TODO: add a file with prerendered responses
			http.Error(w, `{"error":"malformed body"}`, http.StatusUnprocessableEntity)
			log.Printf("WARN: parsing body: %v\n", err)
			return
		}

		ctx := context.WithValue(r.Context(), key, dto)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
