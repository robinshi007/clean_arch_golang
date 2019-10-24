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

func TestUser(t *testing.T) {
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

	uHanlder := handler.NewUserHandler(db)

	server := httptest.NewServer(handler.NewUserRouter(uHanlder))
	defer server.Close()

	e := httpexpect.New(t, server.URL)
	e.GET("/").
		Expect().
		Status(http.StatusOK).JSON().Array().Length().Equal(3)
	e.GET("/1").
		Expect().
		Status(http.StatusOK).JSON().Object().ContainsKey("name").ValueEqual("name", "hi")
}
