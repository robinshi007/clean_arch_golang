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

type AccountUcaseSuite struct {
	suite.Suite
}

func (suite *AccountUcaseSuite) SetupSuite() {
	registry.InitGlobals(WD)
}
func (suite *AccountUcaseSuite) SetupTest() {
	util.MigrationUp(registry.Cfg, WD)
}
func (suite *AccountUcaseSuite) TearDownTest() {
	util.MigrationDown(registry.Cfg, WD)
}

func (suite *AccountUcaseSuite) TestFindAll_0() {
	au := usecase.NewAccountUsecase(
		postgres.NewAccountRepo(),
		presenter.NewAccountPresenter(),
	)
	ctx := context.Background()

	accounts, err := au.FindAll(ctx, &in.FetchAllOptions{})
	util.FailedIf(err)
	expectedCount := 0
	suite.Equal(expectedCount, len(accounts))
}
