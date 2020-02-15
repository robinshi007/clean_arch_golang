package out

import "time"

// Profile -
type Profile struct {
	ID        int64     `json:"id" msgpack:"id"`
	Name      string    `json:"name" msgpack:"name"`
	Email     string    `json:"email" msgpack:"email"`
	CreatedAt time.Time `json:"created_at" msgpack:"created_at"`
	UpdatedAt time.Time `json:"updated_at" msgpack:"updated_at"`
}
