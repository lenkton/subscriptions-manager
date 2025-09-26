package subscriptions

import "net/http"

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) HandleListSubscriptions(w http.ResponseWriter, r *http.Request)  {}
func (s *Service) HandleCreateSubscription(w http.ResponseWriter, r *http.Request) {}
func (s *Service) HandleGetSubscription(w http.ResponseWriter, r *http.Request)    {}
func (s *Service) HandleUpdateSubscription(w http.ResponseWriter, r *http.Request) {}
func (s *Service) HandleDeleteSubscription(w http.ResponseWriter, r *http.Request) {}
