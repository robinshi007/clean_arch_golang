package web

import (
	"net/http"
	"time"

	"clean_arch/infra/database"
)

// NewServer -
func NewServer(conn database.DBM) *http.Server {

	r := NewRouter(conn)
	srv := &http.Server{
		Addr:           ":8005",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	return srv
}
