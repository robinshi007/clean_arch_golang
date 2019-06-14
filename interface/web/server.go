package web

import (
	"net/http"
	"time"

	"github.com/robinshi007/goweb/db"
)

func NewServer(conn *db.Db) *http.Server {

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
