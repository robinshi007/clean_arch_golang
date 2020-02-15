package postgres_test

import (
	"context"

	"github.com/stretchr/testify/suite"

	"clean_arch/adapter/postgres"
	"clean_arch/domain/model"
	"clean_arch/domain/repository"
	"clean_arch/infra/util"
	"clean_arch/registry"
)

type UserRepoSuite struct {
	suite.Suite
}

func (suite *UserRepoSuite) SetupSuite() {
	registry.InitGlobals(WD)
}
func (suite *UserRepoSuite) SetupTest() {
	util.MigrationUp(registry.Cfg, WD)
}
func (suite *UserRepoSuite) TearDownTest() {
	util.MigrationDown(registry.Cfg, WD)
}

func (suite *UserRepoSuite) TestFindAll_0() {
	ur := postgres.NewUserRepo()
	ctx := context.Background()

	users, _ := ur.FindAll(ctx, &repository.ListOptions{})
	expectedCount := 0
	suite.Equal(expectedCount, len(users))
}

func (suite *UserRepoSuite) TestCRUD() {
	ur := postgres.NewUserRepo()
	ctx := context.Background()

	expectedName := "Hello"
	userID, err := ur.Create(ctx, &model.User{Name: expectedName})
	user, err := ur.FindByID(ctx, userID)
	suite.Equal(expectedName, user.Name)

	user2, err := ur.FindByName(ctx, expectedName)
	suite.Equal(expectedName, user2.Name)

	users, err := ur.FindAll(ctx, &repository.ListOptions{})
	expectedCount := 1
	suite.Equal(expectedCount, len(users))

	expectedName2 := "Hello world!"
	user3, err := ur.Update(ctx, &model.User{ID: userID, Name: expectedName2})
	suite.Equal(expectedName2, user3.Name)

	_, err = ur.Create(ctx, &model.User{Name: "Hello Again"})
	users, err = ur.FindAll(ctx, &repository.ListOptions{})
	expectedCount = 2
	suite.Equal(expectedCount, len(users))

	err = ur.Delete(ctx, userID)
	users, err = ur.FindAll(ctx, &repository.ListOptions{})
	util.FailedIf(err)
	expectedCount = 1
	suite.Equal(expectedCount, len(users))
}
