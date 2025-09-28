package subscriptions

import (
	"errors"
	"slices"
)

type Storage struct {
	subscriptions []*Subscription
	lastID        int
}

var ErrSubscriptionNotFound = errors.New("subscription not found")

func NewStorage() *Storage {
	return &Storage{subscriptions: make([]*Subscription, 0)}
}

func (s *Storage) List() []*Subscription {
	return s.subscriptions
}

func (s *Storage) Get(id int) (*Subscription, error) {
	for _, si := range s.subscriptions {
		if si.ID == id {
			return si, nil
		}
	}
	return nil, ErrSubscriptionNotFound
}

// WARN: it modifies the sub: it sets the ID
func (s *Storage) Add(sub *Subscription) int {
	s.lastID++
	sub.ID = s.lastID
	s.subscriptions = append(s.subscriptions, sub)
	return sub.ID
}

func (s *Storage) Update(sub *Subscription) error {
	for _, si := range s.subscriptions {
		if si.ID == sub.ID {
			*si = *sub
			return nil
		}
	}
	return ErrSubscriptionNotFound
}

func (s *Storage) Delete(id int) {
	s.subscriptions = slices.DeleteFunc(s.subscriptions,
		func(si *Subscription) bool { return si.ID == id })
}
