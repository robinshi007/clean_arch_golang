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
	ID        int64     `json:"id"`
	Code      string    `json:"code"`
	URL       string    `json:"url"`
	CreatedAt time.Time `json:"created_at"`
}

// NewRedirect -
func NewRedirect(url string) *Redirect {
	return &Redirect{
		URL: url,
	}
}
