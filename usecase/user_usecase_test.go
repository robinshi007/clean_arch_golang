package usecase_test

import (
	"context"

	"clean_arch/adapter/postgres"
	"clean_arch/adapter/presenter"
	"clean_arch/domain/usecase/in"
	"clean_arch/infra/util"
	"clean_arch/registry"
	"clean_arch/usecase"

	"github.com/stretchr/testify/suite"
)

type UserUcaseSuite struct {
	suite.Suite
}

func (suite *UserUcaseSuite) SetupSuite() {
	registry.InitGlobals(WD)
}
func (suite *UserUcaseSuite) SetupTest() {
	util.MigrationUp(registry.Cfg, WD)
}
func (suite *UserUcaseSuite) TearDownTest() {
	util.MigrationDown(registry.Cfg, WD)
}

func (suite *UserUcaseSuite) TestFindAll_0() {
	uu := usecase.NewUserUsecase(
		postgres.NewUserRepo(),
		presenter.NewUserPresenter(),
	)
	ctx := context.Background()

	users, err := uu.FindAll(ctx, &in.FetchAllOptions{})
	util.FailedIf(err)
	expectedCount := 0
	suite.Equal(expectedCount, len(users))
}
func (suite *UserUcaseSuite) TestCreate() {
	uu := usecase.NewUserUsecase(
		postgres.NewUserRepo(),
		presenter.NewUserPresenter(),
	)
	ctx := context.Background()

	expectedName := "TestName"
	userNewInput := &in.NewUser{
		Name: expectedName,
	}
	// save first one and find one
	userID, err := uu.Create(ctx, userNewInput)
	user, err := uu.FindByID(ctx, &in.FetchUser{
		ID: string(userID),
	})
	util.FailedIf(err)
	suite.Equal(expectedName, user.Name)
}
func (suite *UserUcaseSuite) TestCRUD() {
	uu := usecase.NewUserUsecase(
		postgres.NewUserRepo(),
		presenter.NewUserPresenter(),
	)
	ctx := context.Background()

	expectedName := "TestName"
	userNewInput := &in.NewUser{
		Name: expectedName,
	}
	// save first one and find one
	userID, err := uu.Create(ctx, userNewInput)
	_, err = uu.Create(ctx, &in.NewUser{
		Name: "hello",
	})
	util.FailedIf(err)

	users, err := uu.FindAll(ctx, &in.FetchAllOptions{})
	util.FailedIf(err)
	expectedCount := 2
	suite.Equal(expectedCount, len(users))

	// update one
	expectedNameAgain := "TestNameAgain"
	_, err = uu.Update(ctx, &in.EditUser{
		ID:   string(userID),
		Name: expectedNameAgain,
	})
	user2, err := uu.FindByID(ctx, &in.FetchUser{
		ID: string(userID),
	})
	suite.Equal(expectedNameAgain, user2.Name)

	// delete one
	err = uu.Delete(ctx, &in.FetchUser{
		ID: string(userID),
	})

	users2, err := uu.FindAll(ctx, &in.FetchAllOptions{})
	util.FailedIf(err)
	expectedCount2 := 1
	suite.Equal(expectedCount2, len(users2))
}
