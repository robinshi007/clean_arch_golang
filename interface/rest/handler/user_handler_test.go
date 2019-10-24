package handler_test

// https://github.com/gavv/httpexpect/blob/master/_examples/fruits_test.go

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gavv/httpexpect"

	"clean_arch/infra/config"
	"clean_arch/infra/database"
	"clean_arch/infra/logger"
	"clean_arch/infra/util"
	"clean_arch/interface/rest/handler"
)

func TestUserHandler(t *testing.T) {
	os.Chdir("../../..")
	wd, _ := os.Getwd()

	cfg, err := config.NewConfig(wd)
	util.FailedIf(err)
	fmt.Println("config:", cfg)

	log, err := logger.NewLogger(cfg)
	log.Info("test logger")

	err = database.NewDB(cfg)
	if err != nil {
		log.Info("database err:", err)
	}
	db := database.GetDB()

	// migration up
	util.MigrationDown(cfg, wd)
	util.MigrationUp(cfg, wd)

	uHanlder := handler.NewUserHandler(db)

	server := httptest.NewServer(handler.NewUserRouter(uHanlder))
	defer server.Close()

	e := httpexpect.New(t, server.URL)
	e.GET("/").
		Expect().
		Status(http.StatusOK).JSON().Object().Value("data").Array().Length().Equal(0)

	user1 := map[string]interface{}{
		"name": "Bob",
	}
	user2 := map[string]interface{}{
		"name": "Alice",
	}
	user3 := map[string]interface{}{
		"name": "Ben",
	}

	e.POST("/").WithJSON(user1).
		Expect().
		Status(http.StatusCreated)
	e.POST("/").WithJSON(user2).
		Expect().
		Status(http.StatusCreated)

	e.GET("/").
		Expect().
		Status(http.StatusOK).JSON().Object().Value("data").Array().Length().Equal(2)

	e.GET("/1").
		Expect().
		Status(http.StatusOK).JSON().Object().Value("data").Object().ContainsKey("name").ValueEqual("name", "Bob")

	e.PUT("/1").WithJSON(user3).
		Expect().
		Status(http.StatusOK)

	e.GET("/1").
		Expect().
		Status(http.StatusOK).JSON().Object().Value("data").Object().ContainsKey("name").ValueEqual("name", "Ben")

	e.DELETE("/2").
		Expect().
		Status(http.StatusOK)

	e.GET("/").
		Expect().
		Status(http.StatusOK).JSON().Object().Value("data").Array().Length().Equal(1)

	// migration down
	//util.MigrationDown(cfg, wd)
}
