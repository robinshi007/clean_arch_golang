package model

import "time"

// UserProfile -
type UserProfile struct {
	UID       int64
	Name      string
	FirstName string
	LastName  string
	FullName  string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
