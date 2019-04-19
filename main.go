package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/spf13/viper"

	"github.com/robinshi007/goweb/db"
	uh "github.com/robinshi007/goweb/interface/http"
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

func main() {
	dbHost := viper.GetString(`database.host`)
	dbPort := viper.GetString(`database.port`)
	dbUser := viper.GetString(`database.user`)
	dbPass := viper.GetString(`database.pass`)
	dbName := viper.GetString(`database.name`)

	//conn, err := db.NewDb("localhost", "5432", "postgres", "postgres", "test")
	conn, err := db.NewDb(dbHost, dbPort, dbUser, dbPass, dbName)
	if err != nil {
		panic(err)
	}

	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)
	uHanlder := uh.NewUserHandler(conn)

	r.Route("/", func(rt chi.Router) {
		rt.Mount("/users", uh.NewUserRouter(uHanlder))
	})

	fmt.Println("Server listen at :8005")
	log.Fatal(http.ListenAndServe(":8005", r))

	// c := make(chan os.Signal)
	// signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	// go func() {
	// 	<-c
	// 	os.Exit(-1)
	// }()

}
