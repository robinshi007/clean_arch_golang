package usecase_test

import (
	"context"

	"github.com/stretchr/testify/suite"

	"clean_arch/adapter/postgres"
	"clean_arch/adapter/presenter"
	"clean_arch/domain/usecase/in"
	"clean_arch/infra/util"
	"clean_arch/registry"
	"clean_arch/usecase"
)

type RedirectUcaseSuite struct {
	suite.Suite
}

func (suite *RedirectUcaseSuite) SetupSuite() {
	registry.InitGlobals(WD)
	// clean at first
	util.MigrationDown(registry.Cfg, WD)
}
func (suite *RedirectUcaseSuite) SetupTest() {
	util.MigrationUp(registry.Cfg, WD)
}
func (suite *RedirectUcaseSuite) TearDownTest() {
	util.MigrationDown(registry.Cfg, WD)
}

func (suite *RedirectUcaseSuite) TestFindAll_0() {
	ru := usecase.NewRedirectUsecase(
		postgres.NewRedirectRepo(),
		presenter.NewRedirectPresenter(),
	)
	ctx := context.Background()

	redirects, err := ru.FindAll(ctx, &in.FetchAllOptions{})
	util.FailedIf(err)
	expectedCount := 0
	suite.Equal(expectedCount, len(redirects))
}
func (suite *RedirectUcaseSuite) TestCreate() {
	ru := usecase.NewRedirectUsecase(
		postgres.NewRedirectRepo(),
		presenter.NewRedirectPresenter(),
	)
	au := usecase.NewAccountUsecase(
		postgres.NewAccountRepo(),
		presenter.NewAccountPresenter(),
	)
	ctx := context.Background()

	accountNewInput := &in.NewAccount{
		Name:     "test",
		Email:    "test@test.com",
		Password: "testtest",
	}
	accoutIDString, err := au.Create(ctx, accountNewInput)
	accountID := string(accoutIDString)

	expectedURL := "http://www.test.com"
	redirectNewInput := &in.NewRedirect{
		URL: expectedURL,
		CID: accountID,
	}
	redirectID, err := ru.Create(ctx, redirectNewInput)
	redirect, err := ru.FindByID(ctx, &in.FetchRedirect{
		ID: string(redirectID),
	})
	util.FailedIf(err)
	suite.Equal(expectedURL, redirect.URL)

	redirect2, err := ru.FindByCode(ctx, &in.FetchRedirectByCode{
		Code: redirect.Code,
	})
	util.FailedIf(err)
	suite.Equal(expectedURL, redirect2.URL)

	redirect3, err := ru.FindByURL(ctx, &in.FetchRedirectByURL{
		URL: redirect.URL,
	})
	util.FailedIf(err)
	suite.Equal(expectedURL, redirect3.URL)

	redirects, err := ru.FindAll(ctx, &in.FetchAllOptions{})
	util.FailedIf(err)
	expectedCount := 1
	suite.Equal(expectedCount, len(redirects))
}
func (suite *RedirectUcaseSuite) TestFindOrCreate() {
	ru := usecase.NewRedirectUsecase(
		postgres.NewRedirectRepo(),
		presenter.NewRedirectPresenter(),
	)
	au := usecase.NewAccountUsecase(
		postgres.NewAccountRepo(),
		presenter.NewAccountPresenter(),
	)
	ctx := context.Background()

	accountNewInput := &in.NewAccount{
		Name:     "test",
		Email:    "test@test.com",
		Password: "testtest",
	}
	accoutIDString, err := au.Create(ctx, accountNewInput)
	accountID := string(accoutIDString)

	expectedURL := "http://www.test.com"
	redirectNewInput := &in.FetchOrCreateRedirect{
		URL: expectedURL,
		CID: accountID,
	}
	_, err = ru.FindOrCreate(ctx, redirectNewInput)

	redirects, err := ru.FindAll(ctx, &in.FetchAllOptions{})
	util.FailedIf(err)
	expectedCount := 1
	suite.Equal(expectedCount, len(redirects))
}

//func TestRedirectUsecase(t *testing.T) {
//	registry.InitGlobals(WD)
//	cfg := registry.Cfg
//
//	// migration
//	util.MigrationDown(cfg, WD)
//	util.MigrationUp(cfg, WD)
//
//	ru := usecase.NewRedirectUsecase(
//		postgres.NewRedirectRepo(),
//		presenter.NewRedirectPresenter(),
//	)
//	au := usecase.NewAccountUsecase(
//		postgres.NewAccountRepo(),
//		presenter.NewAccountPresenter(),
//	)
//	ctx := context.Background()
//
//	accountNewInput := &in.NewAccount{
//		Name:     "test",
//		Email:    "test@test.com",
//		Password: "testtest",
//	}
//	accoutIDString, err := au.Create(ctx, accountNewInput)
//	accountID, err := strconv.ParseInt(string(accoutIDString), 10, 64)
//
//	expectedURL := "http://www.test.com"
//	redirectNewInput := &in.NewRedirect{
//		URL: expectedURL,
//		CID: util.Int642String(accountID),
//	}
//	// save first one
//	redirectID, err := ru.Create(ctx, redirectNewInput)
//	fmt.Println("rid", redirectID)
//	redirect, err := ru.FindByID(ctx, &in.FetchRedirect{
//		ID: string(redirectID),
//	})
//	if err != nil {
//		util.FailedIf(err)
//	}
//	if redirect.URL != expectedURL {
//		t.Errorf("RedirectUsecase.FindByID() return redirect with URL %s , expected %s", redirect.URL, expectedURL)
//	}
//	redirect2, err := ru.FindByCode(ctx, &in.FetchRedirectByCode{
//		Code: redirect.Code,
//	})
//	if redirect2.URL != expectedURL {
//		t.Errorf("RedirectUsecase.FindByCode() return redirect with URL %s , expected %s", redirect2.URL, expectedURL)
//	}
//	// save new again
//	expectedURL2 := "http://www.example.com"
//	_, err = ru.Create(ctx, &in.NewRedirect{
//		URL: expectedURL2,
//		CID: util.Int642String(accountID),
//	})
//	redirects, err := ru.FindAll(ctx, &in.FetchAllOptions{})
//	if len(redirects) != 2 {
//		t.Errorf("RedirectUsecase.FindAll() return redirects with length %d , expected %d", len(redirects), 2)
//	}
//	if redirects[0].URL != expectedURL {
//		t.Errorf("RedirectUsecase.FindAll() return first redirect with URL %s , expected %s", redirects[0].URL, expectedURL)
//	}
//	if redirects[1].URL != expectedURL2 {
//		t.Errorf("RedirectUsecase.FindAll() return last redirect with URL %s , expected %s", redirects[1].URL, expectedURL2)
//	}
//
//	if err != nil {
//		t.Errorf("error occurs: %s", err.Error())
//	}
//
//	util.MigrationDown(cfg, WD)
//}
