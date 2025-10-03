package httputil

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strconv"
)

func WithPathID[T any](next http.Handler, pathPartName string, key any) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		idStr := r.PathValue(pathPartName)
		if idStr == "" {
			http.Error(w, `{"error":"not found"}`, http.StatusNotFound)
			log.Printf("DEBUG: parsing path id: empty part: %v\n", pathPartName)
			return
		}
		var id T
		err := parseString(idStr, &id)
		if errors.Is(err, ErrUnknownType) {
			http.Error(w, `{"error":"internal server error"}`, http.StatusInternalServerError)
			log.Printf("ERROR: parsing path id: %v, do not know how to handle type %T\n", err, id)
			return
		}
		var malformedErr *MalformedIDError
		if errors.As(err, &malformedErr) {
			http.Error(w, `{"error":"malformed path"}`, http.StatusUnprocessableEntity)
			log.Printf("DEBUG: parsing path id: %v\n", malformedErr)
			return
		}
		if err != nil {
			http.Error(w, `{"error":"internal server error"}`, http.StatusInternalServerError)
			log.Printf("ERROR: parsing path id: %v\n", err)
			return
		}

		ctx := context.WithValue(r.Context(), key, id)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

var ErrUnknownType = errors.New("unknown type")

// NOTE: maybe the name is not great
type MalformedIDError struct{ error }

func (e *MalformedIDError) Error() string {
	return fmt.Sprintf("invalid id format: %v", e.error)
}

// TODO: add support for string ids
func parseString[T any](s string, t *T) error {
	switch reflect.TypeFor[T]() {
	case reflect.TypeFor[int]():
		i, err := strconv.Atoi(s)
		if err != nil {
			return &MalformedIDError{error: err}
		}
		v := reflect.ValueOf(i)
		// NOTE: this can panic
		reflect.ValueOf(t).Elem().Set(v)
		return nil
	default:
		return ErrUnknownType
	}
}
