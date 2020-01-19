package handler_test

// https://github.com/gavv/httpexpect/blob/master/_examples/fruits_test.go

import (
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/gavv/httpexpect"

	"clean_arch/endpoint/api/handler"
	"clean_arch/infra/util"
	"clean_arch/registry"
)

func TestGraphQLHandlerRedirect(t *testing.T) {
	wd, _ := os.Getwd()
	wd = filepath.Dir(filepath.Dir(filepath.Dir(wd)))

	registry.InitGlobals(wd)
	cfg := registry.Cfg
	db := registry.Db
	defer db.Close()

	// migration down
	util.MigrationDown(cfg, wd)
	util.MigrationUp(cfg, wd)

	gqlHanlder := handler.GraphQLHandler()

	server := httptest.NewServer(gqlHanlder)
	defer server.Close()

	e := httpexpect.New(t, server.URL)

	queryList := map[string]interface{}{
		"operationName": nil,
		"variables":     map[string]interface{}{},
		"query": `{
			redirects {
				id
				code
				url
			}
		}
	`,
	}
	e.POST("/").WithJSON(queryList).
		Expect().
		Status(http.StatusOK).JSON().Object().Value("data").
		Object().Value("redirects").Array().Length().Equal(0)

	mutationCreate := map[string]interface{}{
		"operationName": nil,
		"variables":     map[string]interface{}{},
		"query": `mutation {
			createRedirect(input: {
			url: "http://www.test.com"
			}) {
				id
				code
				url
			}
		}
	`,
	}
	e.POST("/").WithJSON(mutationCreate).
		Expect().
		Status(http.StatusOK).JSON().Object().Value("data").
		Object().Value("createRedirect").Object().ContainsKey("code")

	e.POST("/").WithJSON(queryList).
		Expect().
		Status(http.StatusOK).JSON().Object().Value("data").
		Object().Value("redirects").Array().Length().Equal(1)

}
