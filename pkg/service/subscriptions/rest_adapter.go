package subscriptions

import (
	"errors"
	"log"
	"net/http"

	"github.com/lenkton/subscriptions-manager/pkg/httputil"
)

const SubscriptionIDPathSegmentName = "subscriptionID"

type SubscriptionContextKey struct{}
type SubscriptionIDContextKey struct{}

type RESTAdapter struct {
	service *Service
}

func newRESTAdapter(s *Service) *RESTAdapter {
	return &RESTAdapter{service: s}
}

func (a *RESTAdapter) HandleListSubscriptions(w http.ResponseWriter, r *http.Request) {
	subs, err := a.service.ListSubscriptions()
	if err != nil {
		log.Printf("ERROR: service.ListSubscriptions: %v\n", err)
		http.Error(w, `{"error":"internal server error"}`, http.StatusInternalServerError)
		return
	}

	httputil.EncodeJSON(w, subs, http.StatusOK)
}

func (a *RESTAdapter) HandleCreateSubscription(w http.ResponseWriter, r *http.Request) {
	var handler http.Handler = http.HandlerFunc(a.handleCreateSubscription)
	handler = httputil.WithJSONBody[Subscription](handler, SubscriptionContextKey{})

	handler.ServeHTTP(w, r)
}

func (a *RESTAdapter) HandleGetSubscription(w http.ResponseWriter, r *http.Request) {
	var handler http.Handler = http.HandlerFunc(a.handleGetSubscription)
	handler = httputil.WithPathID[int](handler, SubscriptionIDPathSegmentName, SubscriptionIDContextKey{})

	handler.ServeHTTP(w, r)
}

func (a *RESTAdapter) HandleUpdateSubscription(w http.ResponseWriter, r *http.Request) {}
func (a *RESTAdapter) HandleDeleteSubscription(w http.ResponseWriter, r *http.Request) {}

func (a *RESTAdapter) handleCreateSubscription(w http.ResponseWriter, r *http.Request) {
	dto := r.Context().Value(SubscriptionContextKey{}).(*Subscription)

	err := a.service.CreateSubscription(dto)
	if err != nil {
		// TODO: add a file with prerendered responses
		http.Error(w, `{"error":"internal server error"}`, http.StatusInternalServerError)
		log.Printf("WARN: creating a subscription: %v\n", err)
		return
	}

	httputil.EncodeJSON(w, dto, http.StatusCreated)
}

func (a *RESTAdapter) handleGetSubscription(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(SubscriptionIDContextKey{}).(int)
	log.Printf("DEBUG: got the id: %v\n", id)

	sub, err := a.service.GetSubscription(id)
	if errors.Is(err, ErrSubscriptionNotFound) {
		http.Error(w, `{"error":"not found"}`, http.StatusNotFound)
		log.Printf("DEBUG: sub not found: %v\n", err)
		return
	}
	if err != nil {
		// TODO: add a file with prerendered responses
		http.Error(w, `{"error":"internal server error"}`, http.StatusInternalServerError)
		log.Printf("WARN: getting a subscription: %v\n", err)
		return
	}

	httputil.EncodeJSON(w, sub, http.StatusOK)
}
