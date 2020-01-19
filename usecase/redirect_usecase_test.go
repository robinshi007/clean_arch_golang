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

func TestRedirectUsecase(t *testing.T) {
	registry.InitGlobals(WD)
	cfg := registry.Cfg
	db := registry.Db
	defer db.Close()

	// migration
	util.MigrationDown(cfg, WD)
	util.MigrationUp(cfg, WD)

	ru := usecase.NewRedirectUsecase(
		postgres.NewRedirectRepo(),
		presenter.NewRedirectPresenter(),
	)
	ctx := context.Background()

	expectedURL := "http://www.test.com"
	redirectNewInput := &in.NewRedirect{
		URL: expectedURL,
	}
	// save first one
	redirectID, err := ru.Create(ctx, redirectNewInput)
	redirect, err := ru.FindByID(ctx, &in.FetchRedirect{
		ID: string(redirectID),
	})
	if redirect.URL != expectedURL {
		t.Errorf("RedirectUsecase.FindByID() return redirect with URL %s , expected %s", redirect.URL, expectedURL)
	}
	redirect2, err := ru.FindByCode(ctx, &in.FetchRedirectByCode{
		Code: redirect.Code,
	})
	if redirect2.URL != expectedURL {
		t.Errorf("RedirectUsecase.FindByCode() return redirect with URL %s , expected %s", redirect2.URL, expectedURL)
	}
	// save new again
	expectedURL2 := "http://www.example.com"
	_, err = ru.Create(ctx, &in.NewRedirect{
		URL: expectedURL2,
	})
	redirects, err := ru.FindAll(ctx, &in.FetchAllOptions{})
	if len(redirects) != 2 {
		t.Errorf("RedirectUsecase.FindAll() return redirects with length %d , expected %d", len(redirects), 2)
	}
	if redirects[0].URL != expectedURL {
		t.Errorf("RedirectUsecase.FindAll() return first redirect with URL %s , expected %s", redirects[0].URL, expectedURL)
	}
	if redirects[1].URL != expectedURL2 {
		t.Errorf("RedirectUsecase.FindAll() return last redirect with URL %s , expected %s", redirects[1].URL, expectedURL2)
	}

	if err != nil {
		t.Errorf("error occurs: %s", err.Error())
	}

}
