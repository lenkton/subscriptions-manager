package subscriptions

// NOTE: maybe need to write some DTOs
type Subscription struct {
	ID          int    `json:"id" db:"subscription_id"`
	ServiceName string `json:"service_name" db:"service_name"`
	Price       int    `json:"price" db:"price"`
	// TODO: make it UUID
	UserID string `json:"user_id" db:"user_id"`
	// TODO: intoduce a custom type
	StartDate customDate  `json:"start_date" db:"start_date"`
	EndDate   *customDate `json:"end_date,omitempty" db:"end_date"` // * - for null from db
}
