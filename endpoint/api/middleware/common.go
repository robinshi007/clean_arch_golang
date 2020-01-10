package middleware

import (
	"net/http"
)

// borrowed from https://github.com/justinas/alice

// Middleware -
type Middleware func(http.Handler) http.Handler

// MiddlewareFunc -
type MiddlewareFunc func(http.HandlerFunc) http.HandlerFunc

// Chain -
type Chain struct {
	mws []Middleware
}

// New -
func New(mws ...Middleware) Chain {
	return Chain{mws}
}

// Then -
func (c Chain) Then(last http.Handler) http.Handler {
	if last == nil {
		last = http.DefaultServeMux
	}
	for i := len(c.mws) - 1; i >= 0; i-- {
		last = c.mws[i](last)
	}
	return last
}

// ThenFunc -
func (c Chain) ThenFunc(fn http.HandlerFunc) http.Handler {
	if fn == nil {
		return c.Then(nil)
	}
	return c.Then(fn)
}

// Append -
func (c Chain) Append(mws ...Middleware) Chain {
	newMW := make([]Middleware, 0, len(c.mws)+len(mws))
	newMW = append(newMW, c.mws...)
	newMW = append(newMW, mws...)
	return Chain{newMW}
}
