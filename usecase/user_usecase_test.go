package usecase_test

import (
	"context"
	"testing"

	"clean_arch/adapter/postgres"
	"clean_arch/adapter/presenter"
	"clean_arch/domain/usecase/in"
	"clean_arch/infra/util"
	"clean_arch/registry"
	"clean_arch/usecase"
)

func TestUserUsecase(t *testing.T) {
	registry.InitGlobals(WD)
	cfg := registry.Cfg
	db := registry.Db
	defer db.Close()

	// migration
	util.MigrationDown(cfg, WD)
	util.MigrationUp(cfg, WD)

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
	if user.Name != expectedName {
		t.Errorf("UserUsecase.FindByID() return user with name %s , expected %s", user.Name, expectedName)
	}
	_, err = uu.Create(ctx, &in.NewUser{Name: "Second Name"})
	// update one
	expectedNameAgain := "TestNameAgain"
	_, err = uu.Update(ctx, &in.EditUser{
		Name: expectedNameAgain,
	})
	user2, err := uu.FindByID(ctx, &in.FetchUser{
		ID: string(userID),
	})
	if user2.Name != expectedName {
		t.Errorf("UserUsecase.FindByID() return user with name %s , expected %s", user.Name, expectedName)
	}
	count, err := uu.Count(ctx)
	if count != 2 {
		t.Errorf("UserUsecase.Count() return user count %d , expected %d", count, 2)
	}
	users, err := uu.FindAll(ctx, &in.FetchAllOptions{})
	if len(users) != 2 {
		t.Errorf("UserUsecase.FindAll() return user count %d , expected %d", len(users), 2)
	}
	// delete one
	err = uu.Delete(ctx, &in.FetchUser{
		ID: string(userID),
	})
	// recheck the count
	count2, err := uu.Count(ctx)
	if count2 != 1 {
		t.Errorf("UserUsecase.Count() return user count %d , expected %d", count2, 1)
	}

	if err != nil {
		t.Errorf("error occurs: %s", err.Error())
	}

}
