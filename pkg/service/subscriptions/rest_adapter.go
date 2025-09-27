package subscriptions

import (
	"net/http"

	"github.com/lenkton/subscriptions-manager/pkg/httputil"
)

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

func (a *RESTAdapter) HandleCreateSubscription(w http.ResponseWriter, r *http.Request) {}
func (a *RESTAdapter) HandleGetSubscription(w http.ResponseWriter, r *http.Request)    {}
func (a *RESTAdapter) HandleUpdateSubscription(w http.ResponseWriter, r *http.Request) {}
func (a *RESTAdapter) HandleDeleteSubscription(w http.ResponseWriter, r *http.Request) {}
