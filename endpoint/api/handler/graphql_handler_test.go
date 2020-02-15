package handler_test

// https://github.com/gavv/httpexpect/blob/master/_examples/fruits_test.go

import (
	"context"
	"net/http"
	"net/http/httptest"

	"github.com/gavv/httpexpect"
	"github.com/stretchr/testify/suite"

	"clean_arch/adapter/postgres"
	"clean_arch/domain/model"
	"clean_arch/endpoint/api/handler"
	"clean_arch/infra/util"
	"clean_arch/registry"
)

type GraphQLHandlerSuite struct {
	suite.Suite
}

func (suite *GraphQLHandlerSuite) SetupSuite() {
	registry.InitGlobals(WD)
}
func (suite *GraphQLHandlerSuite) SetupTest() {
	util.MigrationUp(registry.Cfg, WD)

	ar := postgres.NewAccountRepo()
	ctx := context.Background()
	_, err := ar.Create(ctx, &model.UserAccount{
		Name:     "test",
		Email:    "test@test.com",
		Password: "testtest",
	})
	util.FailedIf(err)
}
func (suite *GraphQLHandlerSuite) TearDownTest() {
	util.MigrationDown(registry.Cfg, WD)
}

func (suite *GraphQLHandlerSuite) TestRedirectCRUD() {
	gqlHanlder := handler.GraphQLHandler()
	server := httptest.NewServer(gqlHanlder)
	defer server.Close()

	e := httpexpect.New(suite.T(), server.URL)

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
		Object().Value("createRedirect").Object().ContainsKey("code").ContainsKey("url")

	e.POST("/").WithJSON(queryList).
		Expect().
		Status(http.StatusOK).JSON().Object().Value("data").
		Object().Value("redirects").Array().Length().Equal(1)
}
