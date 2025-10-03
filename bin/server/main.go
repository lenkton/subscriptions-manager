package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/lenkton/subscriptions-manager/pkg/middleware"
	"github.com/lenkton/subscriptions-manager/pkg/service/subscriptions"
)

func main() {
	dbpool, err := connectDB(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("ERROR: cannot connect to db: %v\n", err)
	}
	defer dbpool.Close()
	log.Printf("INFO: successfully connected to database!\n")

	service := subscriptions.NewService(dbpool)
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

// TODO: move it somewhere else
func connectDB(dbURL string) (*pgxpool.Pool, error) {
	pool, err := pgxpool.New(context.Background(), dbURL)
	if err != nil {
		return nil, fmt.Errorf("cannot create pgx pool: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = pool.Ping(ctx)
	if err != nil {
		// NOTE: maybe it is not how it's done
		pool.Close()
		return nil, fmt.Errorf("cannot make a request to db: %v", err)
	}

	return pool, nil
}
