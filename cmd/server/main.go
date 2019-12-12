package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"clean_arch/endpoint/api/server"
	"clean_arch/infra/util"
	"clean_arch/registry"
)

func main() {

	currentPath, _ := os.Getwd()

	registry.InitGlobals(currentPath)
	defer registry.Db.Close()

	// migration up
	fmt.Println("migration database...")
	util.MigrationUp(registry.Cfg, currentPath)

	// print this in dev mode
	if registry.Cfg.Mode == "dev" {
		fmt.Println("config", registry.Cfg)
		registry.Log.Info("Init Logger")
	}

	// server
	errs := make(chan error, 2)
	srv := server.NewServer(registry.Cfg, registry.Db)
	go func() {
		fmt.Println(fmt.Sprintf("server listen at :%s", registry.Cfg.Server.Port))
		errs <- srv.ListenAndServe()
	}()

	// handle signal
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, os.Kill)
		errs <- fmt.Errorf("%s", <-c)
	}()

	fmt.Printf(" %s, shutdown http servers...", <-errs)
	if err := srv.Shutdown(context.Background()); err != nil {
		fmt.Printf("error occured: %v", err)
	} else {
		fmt.Println("done.")
	}
}
