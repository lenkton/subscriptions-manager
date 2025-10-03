package subscriptions

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Storage struct {
	dbpool *pgxpool.Pool
}

var ErrSubscriptionNotFound = errors.New("subscription not found")

func NewStorage(dbpool *pgxpool.Pool) *Storage {
	return &Storage{dbpool: dbpool}
}

func (s *Storage) List() ([]*Subscription, error) {
	// TODO: pagination
	rows, err := s.dbpool.Query(context.Background(), "select * from subscriptions")
	if err != nil {
		return []*Subscription{}, fmt.Errorf("query: %v", err)
	}
	subs, err := pgx.CollectRows(rows, pgx.RowToAddrOfStructByName[Subscription])
	if err != nil {
		return []*Subscription{}, fmt.Errorf("CollectRows: %v", err)
	}
	return subs, nil
}

func (s *Storage) Get(id int) (*Subscription, error) {
	row, err := s.dbpool.Query(context.Background(), "select * from subscriptions where subscription_id=$1", id)
	if err != nil {
		return nil, fmt.Errorf("query: %v", err)
	}
	sub, err := pgx.CollectExactlyOneRow(row, pgx.RowToAddrOfStructByName[Subscription])
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, ErrSubscriptionNotFound
	}
	if err != nil {
		return nil, fmt.Errorf("CollectExactlyOneRow: %v", err)
	}
	return sub, nil
}

// WARN: it modifies the sub: it sets the ID
func (s *Storage) Add(sub *Subscription) (int, error) {
	var id int
	err := s.dbpool.QueryRow(context.Background(),
		"INSERT INTO subscriptions"+
			"(service_name, price, user_id, start_date, end_date)"+
			"VALUES ($1, $2, $3, $4, $5)"+
			"RETURNING subscription_id",
		sub.ServiceName, sub.Price, sub.UserID, sub.StartDate, sub.EndDate,
	).Scan(&id)

	if err != nil {
		return 0, fmt.Errorf("QueryRow+Scan: %v", err)
	}

	sub.ID = id
	return id, nil
}

// TODO: use ErrSubscriptionNotFound
// NOTE: maybe add an id as a separate param?
func (s *Storage) Update(sub *Subscription) error {
	_, err := s.dbpool.Exec(context.Background(),
		"UPDATE subscriptions"+
			"SET service_name = $1, price = $2, user_id = $3, start_date = $4, end_date = $5"+
			"WHERE subscription_id = $6",
		sub.ServiceName, sub.Price, sub.UserID, sub.StartDate, sub.EndDate,
		sub.ID,
	)

	if err != nil {
		return fmt.Errorf("exec: %v", err)
	}

	return nil
}

// TODO: use ErrSubscriptionNotFound
// TODO: updated the signature: now it can return an error
func (s *Storage) Delete(id int) error {
	_, err := s.dbpool.Exec(context.Background(),
		"DELETE FROM subscriptions"+
			"WHERE subscription_id = $1",
		id,
	)

	if err != nil {
		return fmt.Errorf("exec: %v", err)
	}

	return nil
}
