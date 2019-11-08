package handler_test

// https://github.com/gavv/httpexpect/blob/master/_examples/fruits_test.go

import (
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/gavv/httpexpect"

	"clean_arch/infra/util"
	"clean_arch/interface/api/handler"
	"clean_arch/registry"
)

func TestUserHandlerCRUD(t *testing.T) {
	wd, _ := os.Getwd()
	wd = filepath.Dir(filepath.Dir(filepath.Dir(wd)))

	registry.InitGlobals(wd)
	cfg := registry.Cfg
	db := registry.Db
	defer db.Close()

	// migration up
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
	util.MigrationDown(cfg, wd)
}

func TestUserHandlerError(t *testing.T) {
	wd, _ := os.Getwd()
	wd = filepath.Dir(filepath.Dir(filepath.Dir(wd)))

	registry.InitGlobals(wd)
	cfg := registry.Cfg
	db := registry.Db
	defer db.Close()

	// migration up
	util.MigrationDown(cfg, wd)
	util.MigrationUp(cfg, wd)

	uHanlder := handler.NewUserHandler(db)

	server := httptest.NewServer(handler.NewUserRouter(uHanlder))
	defer server.Close()

	e := httpexpect.New(t, server.URL)

	user1 := map[string]interface{}{
		"name": "",
	}
	user2 := map[string]interface{}{
		"age": "12",
	}
	user3 := map[string]interface{}{
		"name": "Bob",
	}
	e.GET("/a1").
		Expect().Status(http.StatusNotFound)
	e.GET("/11").
		Expect().Status(http.StatusConflict)

	e.POST("/").WithJSON(user1).
		Expect().Status(http.StatusBadRequest)
	e.POST("/").WithJSON(user2).
		Expect().Status(http.StatusBadRequest)
	e.POST("/").WithJSON(user3).
		Expect().Status(http.StatusCreated)
	e.POST("/").WithJSON(user3).
		Expect().Status(http.StatusBadRequest)

	e.PUT("/1").WithJSON(user2).
		Expect().Status(http.StatusBadRequest)
	e.PUT("/1").WithJSON(user3).
		Expect().Status(http.StatusInternalServerError)
	e.PUT("/99").WithJSON(user3).
		Expect().Status(http.StatusBadRequest)

	e.DELETE("/a").
		Expect().Status(http.StatusNotFound)
	e.DELETE("/11").
		Expect().Status(http.StatusConflict)

	e.GET("/").
		Expect().
		Status(http.StatusOK).JSON().Object().Value("data").Array().Length().Equal(1)

	// migration down
	util.MigrationDown(cfg, wd)
}
