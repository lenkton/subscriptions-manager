package subscriptions

import (
	"log"
	"net/http"

	"github.com/lenkton/subscriptions-manager/pkg/httputil"
)

type SubscriptionContextKey struct{}

type RESTAdapter struct {
	service *Service
}

func newRESTAdapter(s *Service) *RESTAdapter {
	return &RESTAdapter{service: s}
}

func (a *RESTAdapter) HandleListSubscriptions(w http.ResponseWriter, r *http.Request) {
	subs := a.service.ListSubscriptions()
	httputil.EncodeJSON(w, subs, http.StatusOK)
}

func (a *RESTAdapter) HandleCreateSubscription(w http.ResponseWriter, r *http.Request) {
	var handler http.Handler = http.HandlerFunc(a.handleCreateSubscription)
	handler = httputil.WithJSONBody[Subscription](handler, SubscriptionContextKey{})

	handler.ServeHTTP(w, r)
}

func (a *RESTAdapter) HandleGetSubscription(w http.ResponseWriter, r *http.Request)    {}
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
