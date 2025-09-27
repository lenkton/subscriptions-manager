package subscriptions

type Subscription struct {
	ID          int    `json:"id"`
	ServiceName string `json:"service_name"`
	Price       int    `json:"price"`
	// TODO: make it UUID
	UserID string `json:"user_id"`
	// TODO: intoduce a custom type
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date,omitempty"`
}
