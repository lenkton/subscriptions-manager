package main

import (
	"log"
	"net/http"

	"github.com/lenkton/subscriptions-manager/pkg/middleware"
	"github.com/lenkton/subscriptions-manager/pkg/service/subscriptions"
)

func main() {
	service := subscriptions.NewService()

	mux := http.NewServeMux()
	mux.HandleFunc("GET /subscriptions", service.HandleListSubscriptions)
	mux.HandleFunc("POST /subscriptions", service.HandleCreateSubscription)
	mux.HandleFunc("GET /subscriptions/{subscriptionID}", service.HandleGetSubscription)
	mux.HandleFunc("PUT /subscriptions/{subscriptionID}", service.HandleUpdateSubscription)
	mux.HandleFunc("DELETE /subscriptions/{subscriptionID}", service.HandleDeleteSubscription)

	handler := middleware.Logger(mux)

	// TODO: use env to fetch host/port
	server := http.Server{
		Addr:    ":8080",
		Handler: handler,
	}

	// TODO: graceful stop
	log.Println("INFO: server starting at :8080")
	log.Fatal(server.ListenAndServe())
}
