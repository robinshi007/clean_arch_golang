package model

import "time"

// UserProfile -
type UserProfile struct {
	UID       int64     `db:"uid"`
	Name      string    `db:"name"`
	FirstName string    `db:"first_name"`
	LastName  string    `db:"last_name"`
	FullName  string    `db:"full_name"`
	Email     string    `db:"email"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	DeletedAt time.Time `db:"deleted_at"`
}
