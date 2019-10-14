package model

import (
	"time"
)

// User -
type User struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name" validate:"required"`
	Desc      string    `json:"description"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
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
