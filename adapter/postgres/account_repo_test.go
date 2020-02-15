package postgres_test

import (
	"context"
	"fmt"

	"github.com/stretchr/testify/suite"

	"clean_arch/adapter/postgres"
	"clean_arch/domain/model"
	"clean_arch/domain/repository"
	"clean_arch/infra/util"
	"clean_arch/registry"
)

type AccountRepoSuite struct {
	suite.Suite
}

func (suite *AccountRepoSuite) SetupSuite() {
	registry.InitGlobals(WD)
}
func (suite *AccountRepoSuite) SetupTest() {
	util.MigrationUp(registry.Cfg, WD)
}
func (suite *AccountRepoSuite) TearDownTest() {
	util.MigrationDown(registry.Cfg, WD)
}

func (suite *AccountRepoSuite) TestFindAll_0() {
	ar := postgres.NewAccountRepo()
	ctx := context.Background()

	accounts, _ := ar.FindAll(ctx, &repository.ListOptions{})
	expectedCount := 0
	suite.Equal(expectedCount, len(accounts))
}
func (suite *AccountRepoSuite) TestCreate() {
	ar := postgres.NewAccountRepo()
	ctx := context.Background()

	var account *model.UserAccount
	expectedName := "Hello"
	expectedEmail := "Hello@test.com"
	expectedPass := "World!"
	accountID, err := ar.Create(ctx, &model.UserAccount{Name: expectedName, Email: expectedEmail, Password: expectedPass})
	account, err = ar.FindByID(ctx, accountID)
	util.FailedIf(err)

	suite.Equal(expectedName, account.Name)
	suite.Equal(expectedEmail, account.Email)
	suite.Equal(expectedPass, account.Password)

	account2, err := ar.FindByEmail(ctx, expectedEmail)
	suite.Equal(expectedName, account.Name)
	suite.Equal(expectedEmail, account2.Email)
	suite.Equal(expectedPass, account2.Password)
}

func (suite *AccountRepoSuite) TestUpdate() {
	ar := postgres.NewAccountRepo()
	ctx := context.Background()

	var account *model.UserAccount
	expectedName := "Hello"
	accountID, err := ar.Create(ctx, &model.UserAccount{Name: expectedName, Email: "Hello@test.com", Password: "Pass"})
	expectedName2 := "Great"
	_, err = ar.Update(ctx, &model.UserAccount{UID: accountID, Name: expectedName2})
	account, err = ar.FindByID(ctx, accountID)
	util.FailedIf(err)
	suite.Equal(expectedName2, account.Name)
}

func (suite *AccountRepoSuite) TestUpdatePassword() {
	ar := postgres.NewAccountRepo()
	ctx := context.Background()

	var account *model.UserAccount
	expectedPass := "World!"
	accountID, err := ar.Create(ctx, &model.UserAccount{Name: "Hello", Email: "Hello@test.com", Password: expectedPass})
	expectedPass2 := "World Again!"
	_, err = ar.UpdatePassword(ctx, &model.UserAccount{UID: accountID, Password: expectedPass2})
	account, err = ar.FindByID(ctx, accountID)
	util.FailedIf(err)
	suite.Equal(expectedPass2, account.Password)
}
func (suite *AccountRepoSuite) TestDelete() {
	ar := postgres.NewAccountRepo()
	ctx := context.Background()

	accountID, _ := ar.Create(ctx, &model.UserAccount{Name: "Hello", Email: "Hello@test.com", Password: "Pass"})

	accounts, _ := ar.FindAll(ctx, &repository.ListOptions{})
	expectedCount := 1
	suite.Equal(expectedCount, len(accounts))

	err := ar.Delete(ctx, accountID)
	util.FailedIf(err)

	accounts2, _ := ar.FindAll(ctx, &repository.ListOptions{})
	expectedCount2 := 0
	suite.Equal(expectedCount2, len(accounts2))
}

func (suite *AccountRepoSuite) TestFindAll_2() {
	ar := postgres.NewAccountRepo()
	pr := postgres.NewProfileRepo()
	ctx := context.Background()

	expectedEmail := "Hello"
	expectedPass := "World!"
	expectedEmail2 := "Hi"
	expectedPass2 := "Great!"
	_, err := ar.Create(ctx, &model.UserAccount{Name: expectedEmail, Email: expectedEmail, Password: expectedPass})
	_, err = ar.Create(ctx, &model.UserAccount{Name: expectedEmail2, Email: expectedEmail2, Password: expectedPass2})
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	accounts, err := ar.FindAll(ctx, &repository.ListOptions{})
	expectedCount := 2
	suite.Equal(expectedCount, len(accounts))
	profiles, err := pr.FindAll(ctx, &repository.ListOptions{})
	suite.Equal(expectedCount, len(profiles))
}
