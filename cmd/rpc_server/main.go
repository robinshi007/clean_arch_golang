package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"google.golang.org/grpc"

	"clean_arch/adapter/postgres"
	"clean_arch/adapter/presenter"
	"clean_arch/endpoint/rpc"
	"clean_arch/registry"
	"clean_arch/usecase"
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

	port := "8081"
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("faild to listen: %v", err)
	}
	repo := postgres.NewUserRepo()
	pre := presenter.NewUserPresenter()
	service := usecase.NewUserUsecase(repo, pre)

	server := grpc.NewServer()

	rpc.Apply(server, service)

	// errs
	errs := make(chan error, 2)

	go func() {
		fmt.Printf("start grpc server port: %s\n", port)
		server.Serve(lis)
	}()

	// handle signal
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, os.Kill)
		errs <- fmt.Errorf("%s", <-c)
	}()

	fmt.Printf(" %s, stopping grpc servers...", <-errs)
	server.GracefulStop()
	fmt.Println("done.")
}
