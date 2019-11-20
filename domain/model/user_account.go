package model

import "time"

// UserAccount -
type UserAccount struct {
	UID                    int64
	Name                   string
	Email                  string
	Password               string
	PasswordSalt           string
	PasswordHashArgorithm  string
	PasswordReminderToken  string
	PasswordReminderExpire time.Time
	Status                 int64
	CreatedAt              time.Time
	UpdatedAt              time.Time
	DeletedAt              time.Time
}
