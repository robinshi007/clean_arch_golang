package server

import (
	"fmt"
	"net/http"
	"time"

	"clean_arch/endpoint/api/globals"
	"clean_arch/infra"
)

// NewServer -
func NewServer(cfg *infra.Config, conn infra.DB) *http.Server {

	// init api globals
	globals.InitResponder()

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
