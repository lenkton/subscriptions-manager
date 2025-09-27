package subscriptions

import (
	"encoding/json"
	"log"
	"net/http"
)

type RESTAdapter struct {
	service *Service
}

func newRESTAdapter(s *Service) *RESTAdapter {
	return &RESTAdapter{service: s}
}

func (a *RESTAdapter) HandleListSubscriptions(w http.ResponseWriter, r *http.Request) {
	subs := a.service.ListSubscriptions()

	err := json.NewEncoder(w).Encode(subs)
	if err != nil {
		log.Printf("ERROR: encoding subs: %v\n", err)
	}
}

func (a *RESTAdapter) HandleCreateSubscription(w http.ResponseWriter, r *http.Request) {}
func (a *RESTAdapter) HandleGetSubscription(w http.ResponseWriter, r *http.Request)    {}
func (a *RESTAdapter) HandleUpdateSubscription(w http.ResponseWriter, r *http.Request) {}
func (a *RESTAdapter) HandleDeleteSubscription(w http.ResponseWriter, r *http.Request) {}
