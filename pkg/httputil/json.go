package httputil

import (
	"encoding/json"
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
