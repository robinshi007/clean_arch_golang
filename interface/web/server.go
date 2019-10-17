package web

import (
	"fmt"
	"net/http"
	"time"

	"clean_arch/infra/config"
	"clean_arch/infra/database"
)

// NewServer -
func NewServer(cfg *config.Config, conn database.DB) *http.Server {

	r := NewRouter(conn)
	srv := &http.Server{
		Addr:           fmt.Sprintf(":%s", cfg.Server.Port),
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	return srv
}
