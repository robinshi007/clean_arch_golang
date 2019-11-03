package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"time"

	"google.golang.org/grpc"

	"clean_arch/adapter/postgres"
	"clean_arch/adapter/presenter"
	"clean_arch/infra/config"
	"clean_arch/infra/database"
	"clean_arch/infra/logger"
	"clean_arch/infra/util"
	"clean_arch/interface/rpc"
	"clean_arch/usecase"
)

func main() {
	currentPath, _ := os.Getwd()

	cfg, err := config.NewConfig(currentPath)
	util.FailedIf(err)
	fmt.Println("config:", cfg)

	log, err := logger.NewLogger(cfg)
	log.Info("test logger")

	err = database.NewDB(cfg)

	if err != nil {
		log.Info("database err:", err)
	}
	db := database.GetDB()

	port := "8081"
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("faild to listen: %v", err)
	}
	repo := postgres.NewUserRepo(db)
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
