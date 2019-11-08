package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"time"

	"google.golang.org/grpc"

	"clean_arch/adapter/postgres"
	"clean_arch/adapter/presenter"
	"clean_arch/interface/rpc"
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
	repo := postgres.NewUserRepo(registry.Db)
	pre := presenter.NewUserPresenter()
	service := usecase.NewUserUseCase(repo, pre, time.Second)

	server := grpc.NewServer()

	rpc.Apply(server, service)

	go func() {
		fmt.Printf("start grpc server port: %s", port)
		server.Serve(lis)
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	fmt.Println("stopping grpc server...")
	server.GracefulStop()
}
