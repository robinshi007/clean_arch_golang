package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	//"time"

	"github.com/spf13/viper"

	"github.com/robinshi007/goweb/db"
	"github.com/robinshi007/goweb/interface/web"
)

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	if viper.GetBool(`debug`) {
		fmt.Println("Service is running inon DEBUG mode")
	}
}
func failedIf(err error) {
	if err != nil {
		panic(err)
	}
}
func main() {
	dbHost := viper.GetString(`database.host`)
	dbPort := viper.GetString(`database.port`)
	dbUser := viper.GetString(`database.user`)
	dbPass := viper.GetString(`database.pass`)
	dbName := viper.GetString(`database.name`)

	//conn, err := db.NewDb("localhost", "5432", "postgres", "postgres", "test")
	conn, err := db.NewDb(dbHost, dbPort, dbUser, dbPass, dbName)
	failedIf(err)
	// server
	srv := web.NewServer(conn)
	go func() {
		fmt.Println("Server listen at :8005")
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			fmt.Printf("Http Server ListenAndServe: %v", err)
		}
	}()

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
