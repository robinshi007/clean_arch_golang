package model

import (
	"time"
)

type User struct {
	Id        int64     `json:"id"`
	Name      string    `json:"name"`
	Desc      string    `json:"description"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}
