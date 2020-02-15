// foreign key https://gist.github.com/anti1869/84b994692c1b0b2de58446cba328026d
package model

import (
	"errors"
	"time"
)

var (
	// ErrRedirectNotFound -
	ErrRedirectNotFound = errors.New("Redirect Not Found")
	// ErrRedirectInvalid -
	ErrRedirectInvalid = errors.New("Redirect Invalid")
)

// Redirect -
type Redirect struct {
	ID        int64       `json:"id" db:"id"`
	Code      string      `json:"code" db:"code"`
	URL       string      `json:"url" db:"url"`
	CreatedBy UserProfile `json:"created_by" db:"created_by"`
	CreatedAt time.Time   `json:"created_at" db:"created_at"`
}

// NewRedirect -
func NewRedirect(url string) *Redirect {
	return &Redirect{
		URL: url,
	}
}
