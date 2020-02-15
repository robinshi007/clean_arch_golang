package handler_test

// https://github.com/gavv/httpexpect/blob/master/_examples/fruits_test.go

import (
	"net/http"
	"net/http/httptest"

	"github.com/gavv/httpexpect"
	"github.com/stretchr/testify/suite"

	"clean_arch/endpoint/api/handler"
	"clean_arch/infra/util"
	"clean_arch/registry"
)

type UserHandlerSuite struct {
	suite.Suite
}

func (suite *UserHandlerSuite) SetupSuite() {
	registry.InitGlobals(WD)
}
func (suite *UserHandlerSuite) SetupTest() {
	util.MigrationUp(registry.Cfg, WD)
}
func (suite *UserHandlerSuite) TearDownTest() {
	util.MigrationDown(registry.Cfg, WD)
}

func (suite *UserHandlerSuite) TestCRUD() {
	uHanlder := handler.NewUserHandler()
	server := httptest.NewServer(handler.NewUserRouter(uHanlder))
	defer server.Close()

	e := httpexpect.New(suite.T(), server.URL)

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
	e.GET("/Ben/by_name").
		Expect().
		Status(http.StatusOK).JSON().Object().Value("data").Object().ContainsKey("id").ValueEqual("id", 1)

	e.DELETE("/2").
		Expect().
		Status(http.StatusOK)

	e.GET("/").
		Expect().
		Status(http.StatusOK).JSON().Object().Value("data").Array().Length().Equal(1)
}

func (suite *UserHandlerSuite) TestError() {
	uHanlder := handler.NewUserHandler()
	server := httptest.NewServer(handler.NewUserRouter(uHanlder))
	defer server.Close()

	e := httpexpect.New(suite.T(), server.URL)

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
		Expect().Status(http.StatusNotFound)

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
		Expect().Status(http.StatusNotModified)
	e.PUT("/a9").WithJSON(user3).
		Expect().Status(http.StatusNotFound)
	e.PUT("/99").WithJSON(user3).
		Expect().Status(http.StatusNotFound)

	e.DELETE("/a").
		Expect().Status(http.StatusNotFound)
	e.DELETE("/11").
		Expect().Status(http.StatusNotFound)

	e.GET("/").
		Expect().
		Status(http.StatusOK).JSON().Object().Value("data").Array().Length().Equal(1)
}
