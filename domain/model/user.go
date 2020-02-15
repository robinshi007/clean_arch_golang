package model

import (
	"time"
)

// User -
type User struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at" db:"deleted_at"`
}

// NewUser -
func NewUser(name string) *User {
	return &User{
		Name: name,
	}
}

// GetID -
func (u *User) GetID() int64 {
	return u.ID
}

// GetName -
func (u *User) GetName() string {
	return u.Name
}
