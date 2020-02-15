package model

import "time"

// UserAccount -
type UserAccount struct {
	UID                    int64  `db:"uid"`
	Name                   string `db:"name"`
	Email                  string `db:"email"`
	Password               string `db:"password"`
	PasswordSalt           string
	PasswordHashArgorithm  string
	PasswordReminderToken  string
	PasswordReminderExpire time.Time
	Status                 int64
	CreatedAt              time.Time `db:"created_at"`
	UpdatedAt              time.Time `db:"updated_at"`
	DeletedAt              time.Time `db:"deleted_at"`
}
