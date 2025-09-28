package subscriptions

type Service struct {
	storage *Storage
}

func NewService() *Service {
	return &Service{
		storage: NewStorage(),
	}
}

func (s *Service) RESTAdapter() *RESTAdapter {
	return newRESTAdapter(s)
}

func (s *Service) ListSubscriptions() []*Subscription {
	return s.storage.List()
}

// NOTE: maybe in the future it will be returning an error
// WARN: it modifies the passed sub: it sets up the ID
func (s *Service) CreateSubscription(sub *Subscription) error {
	s.storage.Add(sub)
	return nil
}
