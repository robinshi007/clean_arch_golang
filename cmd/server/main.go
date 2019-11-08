package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"clean_arch/interface/api/server"
	"clean_arch/registry"
)

func main() {

	currentPath, _ := os.Getwd()

	registry.InitGlobals(currentPath)
	defer registry.Db.Close()

	// print this in dev mode
	if registry.Cfg.Mode == "dev" {
		fmt.Println("config", registry.Cfg)
		registry.Log.Info("Init Logger")
	}

	// server
	srv := server.NewServer(registry.Cfg, registry.Db)
	go func() {
		fmt.Println(fmt.Sprintf("Server listen at :%s", registry.Cfg.Server.Port))
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
