package model

import (
	"time"
)

type User struct {
	Id        int64     `json:"id"`
	Name      string    `json:"name" validate:"required"`
	Desc      string    `json:"description"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

func NewUser(name string) *User {
	return &User{
		Name: name,
	}
}

func (u *User) GetID() int64 {
	return u.Id
}
func (u *User) GetName() string {
	return u.Name
}
