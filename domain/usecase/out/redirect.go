package out

import "time"

// Redirect -
type Redirect struct {
	ID        int64     `json:"id" msgpack:"id"`
	Code      string    `json:"code" msgpack:"code"`
	URL       string    `json:"url" msgpack:"url"`
	CreatedAt time.Time `json:"created_at" msgpack:"created_at"`
}
