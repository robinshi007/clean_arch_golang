package postgres_test

import (
	"context"

	"github.com/stretchr/testify/suite"
	"github.com/teris-io/shortid"

	"clean_arch/adapter/postgres"
	"clean_arch/domain/model"
	"clean_arch/domain/repository"
	"clean_arch/infra/util"
	"clean_arch/registry"
)

type RedirectRepoSuite struct {
	suite.Suite
}

func (suite *RedirectRepoSuite) SetupSuite() {
	registry.InitGlobals(WD)
}
func (suite *RedirectRepoSuite) SetupTest() {
	util.MigrationUp(registry.Cfg, WD)
}
func (suite *RedirectRepoSuite) TearDownTest() {
	util.MigrationDown(registry.Cfg, WD)
}

func (suite *RedirectRepoSuite) TestFindAll() {
	rr := postgres.NewRedirectRepo()
	ctx := context.Background()

	redirects, _ := rr.FindAll(ctx, &repository.ListOptions{})
	expectedCount := 0
	suite.Equal(expectedCount, len(redirects))
}
func (suite *RedirectRepoSuite) TestCreate() {
	rr := postgres.NewRedirectRepo()
	ar := postgres.NewAccountRepo()
	ctx := context.Background()

	accountID, err := ar.Create(ctx, &model.UserAccount{
		Name:     "test",
		Email:    "test@test.com",
		Password: "testtest",
	})
	util.FailedIf(err)

	expectedURL := "http://www.test.com"
	code := shortid.MustGenerate()
	redirectID, err := rr.Create(ctx, &model.Redirect{
		Code:      code,
		URL:       expectedURL,
		CreatedBy: model.UserProfile{UID: accountID},
	})

	redirects, err := rr.FindAll(ctx, &repository.ListOptions{})
	expectedCount := 1
	suite.Equal(expectedCount, len(redirects))

	redirect, err := rr.FindByCode(ctx, code)
	suite.Equal(expectedURL, redirect.URL)

	redirect2, err := rr.FindByID(ctx, redirectID)
	suite.Equal(expectedURL, redirect2.URL)

}
func (suite *RedirectRepoSuite) TestDelete() {
	rr := postgres.NewRedirectRepo()
	ar := postgres.NewAccountRepo()
	ctx := context.Background()

	accountID, err := ar.Create(ctx, &model.UserAccount{
		Name:     "test",
		Email:    "test@test.com",
		Password: "testtest",
	})
	util.FailedIf(err)

	expectedURL := "http://www.test.com"
	code := shortid.MustGenerate()
	_, err = rr.Create(ctx, &model.Redirect{
		Code:      code,
		URL:       expectedURL,
		CreatedBy: model.UserProfile{UID: accountID},
	})

	expectedURL2 := "http://www.example.com"
	code2 := shortid.MustGenerate()
	redirectID, err := rr.Create(ctx, &model.Redirect{
		Code:      code2,
		URL:       expectedURL2,
		CreatedBy: model.UserProfile{UID: accountID},
	})

	redirects, err := rr.FindAll(ctx, &repository.ListOptions{
		Query: "",
		LimitOffset: &repository.LimitOffset{
			Limit:  5,
			Offset: 0,
		},
	})
	util.FailedIf(err)
	suite.Equal(2, len(redirects))

	err = rr.Delete(ctx, int64(redirectID))
	util.FailedIf(err)
	redirects2, err := rr.FindAll(ctx, &repository.ListOptions{})
	suite.Equal(1, len(redirects2))
}
