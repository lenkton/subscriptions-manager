package subscriptions

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) RESTAdapter() *RESTAdapter {
	return newRESTAdapter()
}
