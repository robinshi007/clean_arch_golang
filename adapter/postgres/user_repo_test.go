package postgres_test

import (
	"context"
	"testing"

	"clean_arch/adapter/postgres"
	"clean_arch/domain/model"
	"clean_arch/domain/repository"
	"clean_arch/infra/util"
	"clean_arch/registry"
)

func TestUserCRUD(t *testing.T) {
	registry.InitGlobals(WD)
	cfg := registry.Cfg
	db := registry.Db
	defer db.Close()

	// migration up

	util.MigrationDown(cfg, WD)
	util.MigrationUp(cfg, WD)

	ur := postgres.NewUserRepo()
	ctx := context.Background()

	users, err := ur.GetAll(ctx, &repository.UserListOptions{})
	expectedCount := 0
	if len(users) != expectedCount {
		t.Errorf("UserRepo.GetAll() return %d user, expected %d", len(users), expectedCount)
	}

	var user *model.User
	expectedName := "Hello"
	userID, err := ur.Create(ctx, &model.User{Name: expectedName})
	user, err = ur.GetByID(ctx, userID)
	if user.Name != expectedName {
		t.Errorf("UserRepo.GetByID() return user with name %s , expected %s", user.Name, expectedName)
	}
	user, err = ur.GetByName(ctx, expectedName)
	if user.Name != expectedName {
		t.Errorf("UserRepo.GetByName() return user with name %s , expected %s", user.Name, expectedName)
	}

	users, err = ur.GetAll(ctx, &repository.UserListOptions{})
	expectedCount = 1
	if len(users) != expectedCount {
		t.Errorf("UserRepo.GetAll() return %d user, expected %d", len(users), expectedCount)
	}

	expectedName = "Hello world!"
	user, err = ur.Update(ctx, &model.User{Name: expectedName})
	if user.Name != expectedName {
		t.Errorf("UserRepo.Update() return user with name %s , expected %s", user.Name, expectedName)
	}
	_, err = ur.Create(ctx, &model.User{Name: "Hello Again"})
	users, err = ur.GetAll(ctx, &repository.UserListOptions{})
	expectedCount = 2
	if len(users) != expectedCount {
		t.Errorf("UserRepo.GetAll() return %d user, expected %d", len(users), expectedCount)
	}
	err = ur.Delete(ctx, userID)
	users, err = ur.GetAll(ctx, &repository.UserListOptions{})
	expectedCount = 1
	if len(users) != expectedCount {
		t.Errorf("UserRepo.GetAll() return %d user, expected %d", len(users), expectedCount)
	}

	if err != nil {
		t.Errorf("error occurs: %s", err.Error())
	}

	util.MigrationDown(cfg, WD)
}
