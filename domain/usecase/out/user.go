package out

import "time"

// ID -
type ID string

// User -
type User struct {
	ID        int64     `json:"id" xml:"id"`
	Name      string    `json:"name" xml:"name"`
	CreatedAt time.Time `json:"created_at" xml:"created_at"`
	UpdatedAt time.Time `json:"updated_at" xml:"updated_at"`
}
