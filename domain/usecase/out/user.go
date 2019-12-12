package out

import "time"

// ID -
type ID string

// User -
type User struct {
	ID        int64     `json:"id" msgpack:"id"`
	Name      string    `json:"name" msgpack:"name"`
	CreatedAt time.Time `json:"created_at" msgpack:"created_at"`
	UpdatedAt time.Time `json:"updated_at" msgpack:"updated_at"`
}
