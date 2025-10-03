package subscriptions

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

// TODO: move somewhere else
type customDate struct {
	time.Time
}

func (t *customDate) UnmarshalJSON(b []byte) error {
	time, err := time.Parse(`"01-2006"`, string(b))
	t.Time = time
	return err
}

func (t *customDate) MarshalJSON() ([]byte, error) {
	return []byte(t.Format(`"01-2006"`)), nil
}

func (dst *customDate) ScanDate(v pgtype.Date) error {
	dst.Time = v.Time
	return nil
}

func (src customDate) DateValue() (pgtype.Date, error) {
	date := pgtype.Date{Valid: true, Time: src.Time}
	return date, nil
}
