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
