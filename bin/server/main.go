package main

import (
	"log"
	"net/http"

	"github.com/lenkton/subscriptions-manager/pkg/middleware"
	"github.com/lenkton/subscriptions-manager/pkg/service/subscriptions"
)

func main() {
	service := subscriptions.NewService()
	adapter := service.RESTAdapter()

	mux := http.NewServeMux()
	mux.HandleFunc("GET /subscriptions", adapter.HandleListSubscriptions)
	mux.HandleFunc("POST /subscriptions", adapter.HandleCreateSubscription)
	mux.HandleFunc("GET /subscriptions/{subscriptionID}", adapter.HandleGetSubscription)
	mux.HandleFunc("PUT /subscriptions/{subscriptionID}", adapter.HandleUpdateSubscription)
	mux.HandleFunc("DELETE /subscriptions/{subscriptionID}", adapter.HandleDeleteSubscription)

	// TODO: mark content type as json
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
