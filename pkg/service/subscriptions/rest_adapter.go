package subscriptions

import "net/http"

type RESTAdapter struct {
}

func newRESTAdapter() *RESTAdapter {
	return &RESTAdapter{}
}

func (a *RESTAdapter) HandleListSubscriptions(w http.ResponseWriter, r *http.Request)  {}
func (a *RESTAdapter) HandleCreateSubscription(w http.ResponseWriter, r *http.Request) {}
func (a *RESTAdapter) HandleGetSubscription(w http.ResponseWriter, r *http.Request)    {}
func (a *RESTAdapter) HandleUpdateSubscription(w http.ResponseWriter, r *http.Request) {}
func (a *RESTAdapter) HandleDeleteSubscription(w http.ResponseWriter, r *http.Request) {}
