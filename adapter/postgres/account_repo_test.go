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

func TestAccountCRUD(t *testing.T) {
	registry.InitGlobals(WD)
	cfg := registry.Cfg
	db := registry.Db
	defer db.Close()

	// migration up
	util.MigrationDown(cfg, WD)
	util.MigrationUp(cfg, WD)

	ar := postgres.NewAccountRepo()
	pr := postgres.NewProfileRepo()
	ctx := context.Background()

	accounts, err := ar.GetAll(ctx, &repository.AccountListOptions{})
	expectedCount := 0
	if len(accounts) != expectedCount {
		t.Errorf("AccountRepo.GetAll() return %d account, expected %d", len(accounts), expectedCount)
	}

	var account *model.UserAccount
	expectedEmail := "Hello"
	expectedPass := "World!"
	accountID, err := ar.Create(ctx, &model.UserAccount{Email: expectedEmail, Password: expectedPass})
	account, err = ar.GetByID(ctx, accountID)
	if account.Email != expectedEmail {
		t.Errorf("AccountRepo.GetByID() return account with email %s , expected %s", account.Email, expectedEmail)
	}
	if account.Password != expectedPass {
		t.Errorf("AccountRepo.GetByID() return account with password %s , expected %s", account.Password, expectedPass)
	}
	account, err = ar.GetByEmail(ctx, expectedEmail)
	if account.Email != expectedEmail {
		t.Errorf("AccountRepo.GetByName() return account with name %s , expected %s", account.Email, expectedEmail)
	}
	accounts, err = ar.GetAll(ctx, &repository.AccountListOptions{})
	expectedCount = 1
	if len(accounts) != expectedCount {
		t.Errorf("AccountRepo.GetAll() return %d account, expected %d", len(accounts), expectedCount)
	}
	profiles, err := pr.GetAll(ctx, &repository.ProfileListOptions{})
	if len(profiles) != expectedCount {
		t.Errorf("ProfileRepo.GetAll() return %d account, expected %d", len(profiles), expectedCount)
	}

	expectedPass = "Hello world!"
	account, err = ar.Update(ctx, &model.UserAccount{UID: accountID, Password: expectedPass})
	if account.Password != expectedPass {
		t.Errorf("AccountRepo.Update() return account with password %s , expected %s", account.Password, expectedPass)
	}
	_, err = ar.Create(ctx, &model.UserAccount{Email: "Hello Again", Password: "pass"})
	accounts, err = ar.GetAll(ctx, &repository.AccountListOptions{})
	expectedCount = 2
	if len(accounts) != expectedCount {
		t.Errorf("AccountRepo.GetAll() return %d account, expected %d", len(accounts), expectedCount)
	}

	err = ar.Delete(ctx, accountID)
	accounts, err = ar.GetAll(ctx, &repository.AccountListOptions{})
	expectedCount = 1
	if len(accounts) != expectedCount {
		t.Errorf("AccountRepo.GetAll() return %d account, expected %d", len(accounts), expectedCount)
	}
	profiles, err = pr.GetAll(ctx, &repository.ProfileListOptions{})
	if len(profiles) != expectedCount {
		t.Errorf("ProfileRepo.GetAll() return %d account, expected %d", len(profiles), expectedCount)
	}
	if err != nil {
		t.Errorf("error occurs: %s", err.Error())
	}

	util.MigrationDown(cfg, WD)
}
