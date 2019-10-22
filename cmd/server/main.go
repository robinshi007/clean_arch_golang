package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"clean_arch/adapter/config"
	"clean_arch/adapter/database"
	"clean_arch/adapter/logger"
	"clean_arch/interface/rest"
	"clean_arch/pkg/util"
)

func main() {

	cfg, err := config.NewConfig()
	util.FailedIf(err)
	fmt.Println("config:", cfg)

	log, err := logger.NewLogger(cfg)
	log.Info("test logger")

	err = database.NewDB(cfg)
	if err != nil {
		log.Info("database err:", err)
	}
	db := database.GetDB()

	// server
	srv := rest.NewServer(cfg, db)
	go func() {
		fmt.Println(fmt.Sprintf("Server listen at :%s", cfg.Server.Port))
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			fmt.Printf("Http Server ListenAndServe: %v", err)
		}
	}()
	//
	// handle signal
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	fmt.Println(<-quit)

	fmt.Print("Stopping http server...")
	if err := srv.Shutdown(context.Background()); err != nil {
		fmt.Printf("Http server Shutdown: %v", err)
	} else {
		fmt.Println("Done.")
	}
}
