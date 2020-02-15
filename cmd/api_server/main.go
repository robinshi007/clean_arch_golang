package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/signal"

	"github.com/casbin/casbin/v2"
	"github.com/teris-io/shortid"

	"clean_arch/adapter/postgres"
	"clean_arch/adapter/presenter"
	"clean_arch/domain/model"
	"clean_arch/domain/usecase/in"
	"clean_arch/endpoint/api/server"
	"clean_arch/infra/util"
	"clean_arch/pkg/casbinsqlx"
	"clean_arch/registry"
	ctn "clean_arch/usecase"
)

func main() {

	currentPath, _ := os.Getwd()

	registry.InitGlobals(currentPath)
	defer registry.Db.Close()

	// migrate schema up
	fmt.Printf("PQSQL migration database...")
	util.MigrationUp(registry.Cfg, currentPath)
	fmt.Println("done")

	// migrate init data
	{
		repo := postgres.NewAccountRepo()
		pre := presenter.NewAccountPresenter()
		uc := ctn.NewAccountUsecase(repo, pre)
		_, err := uc.FindByEmail(context.Background(), &in.FetchAccountByEmail{
			Email: "admin@test.com",
		})
		// if not found super admin user, create a new one
		if errors.Is(err, model.ErrEntityNotFound) {
			fmt.Printf("=> create super user 'admin'...")
			password := shortid.MustGenerate()
			fmt.Printf("with '%s'...", password)
			uc.Create(context.Background(), &in.NewAccount{
				Name:     "admin",
				Email:    "admin@test.com",
				Password: password,
			})
			fmt.Println("Done")
		}
		_, err = uc.FindByEmail(context.Background(), &in.FetchAccountByEmail{
			Email: "test@test.com",
		})
		// if not found super admin user, create a new one
		if errors.Is(err, model.ErrEntityNotFound) {
			fmt.Printf("=> create test user 'test'...")
			uc.Create(context.Background(), &in.NewAccount{
				Name:     "test",
				Email:    "test@test.com",
				Password: "testtest",
			})
			fmt.Println("Done")
		}

		rulesRepo := postgres.NewCasbinRuleRepo()
		rulesCount, err := rulesRepo.Count(context.Background())
		util.FailedIf(err)

		if rulesCount == 0 {
			//e, err := casbin.NewEnforcer("./examples/rbac_model.conf", "examples/rbac_policy.csv")
			e, err := casbin.NewEnforcer("./config/casbin_rbac_model.conf", "./config/casbin_rbac_policy.csv")
			if err != nil {
				panic(err)
			}
			a := casbinsqlx.NewAdapterByDB(registry.Db)
			err = a.SavePolicy(e.GetModel())
			if err != nil {
				panic(err)
			}
		}
	}

	// print this in dev mode
	if registry.Cfg.Mode == "dev" {
		fmt.Println("config", registry.Cfg)
		registry.Log.Info("Init Logger")
	}

	// server
	errs := make(chan error, 2)
	srv := server.NewServer(registry.Cfg)
	go func() {
		fmt.Println(fmt.Sprintf("server listen at :%s\n", registry.Cfg.Server.Port))
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
