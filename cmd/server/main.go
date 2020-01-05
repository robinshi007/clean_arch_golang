package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/signal"
	"time"

	"clean_arch/adapter/postgres"
	"clean_arch/adapter/presenter"
	"clean_arch/domain/model"
	"clean_arch/domain/usecase/in"
	"clean_arch/endpoint/api/server"
	"clean_arch/infra/util"
	"clean_arch/registry"
	ctn "clean_arch/usecase"
)

func main() {

	currentPath, _ := os.Getwd()

	registry.InitGlobals(currentPath)
	defer registry.Db.Close()

	// migrate schema up
	fmt.Println("migration database...")
	util.MigrationUp(registry.Cfg, currentPath)

	// migrate init data
	{
		repo := postgres.NewAccountRepo()
		pre := presenter.NewAccountPresenter()
		uc := ctn.NewAccountUseCase(repo, pre, 2*time.Second)
		_, err := uc.GetByEmail(context.Background(), &in.FetchAccountByEmail{
			Email: "admin@test.com",
		})
		// if not found super admin user, create a new one
		if errors.Is(err, model.ErrEntityNotFound) {
			fmt.Printf("=> create super user 'admin'...")
			uc.Create(context.Background(), &in.NewAccount{
				Name:     "admin",
				Email:    "admin@test.com",
				Password: "password", // will change in production env
			})
			fmt.Println("Done")
		}
	}

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
