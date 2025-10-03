package subscriptions

import "github.com/jackc/pgx/v5/pgxpool"

type Service struct {
	storage *Storage
}

func NewService(connectionPool *pgxpool.Pool) *Service {
	return &Service{
		storage: NewStorage(connectionPool),
	}
}

func (s *Service) RESTAdapter() *RESTAdapter {
	return newRESTAdapter(s)
}

func (s *Service) ListSubscriptions() ([]*Subscription, error) {
	return s.storage.List()
}

// WARN: it modifies the passed sub: it sets up the ID
func (s *Service) CreateSubscription(sub *Subscription) error {
	_, err := s.storage.Add(sub)
	return err
}

func (s *Service) GetSubscription(id int) (*Subscription, error) {
	return s.storage.Get(id)
}
