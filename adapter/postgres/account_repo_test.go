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

	// migration down
	util.MigrationDown(cfg, WD)
	util.MigrationUp(cfg, WD)

	ar := postgres.NewAccountRepo()
	pr := postgres.NewProfileRepo()
	ctx := context.Background()

	accounts, err := ar.FindAll(ctx, &repository.ListOptions{})
	expectedCount := 0
	if len(accounts) != expectedCount {
		t.Errorf("AccountRepo.FindAll() return %d account, expected %d", len(accounts), expectedCount)
	}

	var account *model.UserAccount
	expectedEmail := "Hello"
	expectedPass := "World!"
	accountID, err := ar.Create(ctx, &model.UserAccount{Email: expectedEmail, Password: expectedPass})
	account, err = ar.FindByID(ctx, accountID)
	if account.Email != expectedEmail {
		t.Errorf("AccountRepo.FindByID() return account with email %s , expected %s", account.Email, expectedEmail)
	}
	if account.Password != expectedPass {
		t.Errorf("AccountRepo.FindByID() return account with password %s , expected %s", account.Password, expectedPass)
	}
	account, err = ar.FindByEmail(ctx, expectedEmail)
	if account.Email != expectedEmail {
		t.Errorf("AccountRepo.FindByName() return account with name %s , expected %s", account.Email, expectedEmail)
	}
	accounts, err = ar.FindAll(ctx, &repository.ListOptions{})
	expectedCount = 1
	if len(accounts) != expectedCount {
		t.Errorf("AccountRepo.FindAll() return %d account, expected %d", len(accounts), expectedCount)
	}
	profiles, err := pr.FindAll(ctx, &repository.ListOptions{})
	if len(profiles) != expectedCount {
		t.Errorf("ProfileRepo.FindAll() return %d account, expected %d", len(profiles), expectedCount)
	}

	expectedName := "Great!"
	_, err = ar.Update(ctx, &model.UserAccount{UID: accountID, Name: expectedName})
	accountNameNew, err := ar.FindByID(ctx, accountID)
	if accountNameNew.Name != expectedName {
		t.Errorf("AccountRepo.Update() return account with name %s , expected %s", accountNameNew.Name, expectedName)
	}

	expectedPassNew := "Hello world!"
	_, err = ar.UpdatePassword(ctx, &model.UserAccount{UID: accountID, Password: expectedPassNew})
	accountPassNew, err := ar.FindByID(ctx, accountID)
	if accountPassNew.Password != expectedPassNew {
		t.Errorf("AccountRepo.Update() return account with password %s , expected %s", accountPassNew.Password, expectedPassNew)
	}

	_, err = ar.Create(ctx, &model.UserAccount{Email: "Hello Again", Password: "pass"})
	accounts, err = ar.FindAll(ctx, &repository.ListOptions{})
	expectedCount = 2
	if len(accounts) != expectedCount {
		t.Errorf("AccountRepo.FindAll() return %d account, expected %d", len(accounts), expectedCount)
	}

	err = ar.Delete(ctx, accountID)
	accounts, err = ar.FindAll(ctx, &repository.ListOptions{})
	expectedCount = 1
	if len(accounts) != expectedCount {
		t.Errorf("AccountRepo.FindAll() return %d account, expected %d", len(accounts), expectedCount)
	}
	profiles, err = pr.FindAll(ctx, &repository.ListOptions{})
	if len(profiles) != expectedCount {
		t.Errorf("ProfileRepo.FindAll() return %d account, expected %d", len(profiles), expectedCount)
	}
	if err != nil {
		t.Errorf("error occurs: %s", err.Error())
	}

}
