package postgres_test

import (
	"context"
	"testing"

	"github.com/teris-io/shortid"

	"clean_arch/adapter/postgres"
	"clean_arch/domain/model"
	"clean_arch/domain/repository"
	"clean_arch/infra/util"
	"clean_arch/registry"
)

func TestRedirectCRUD(t *testing.T) {
	registry.InitGlobals(WD)
	cfg := registry.Cfg
	db := registry.Db
	defer db.Close()

	// migration up
	util.MigrationDown(cfg, WD)
	util.MigrationUp(cfg, WD)

	rr := postgres.NewRedirectRepo()
	ctx := context.Background()

	var redirect *model.Redirect
	expectedURL := "http://www.test.com"
	code := shortid.MustGenerate()
	redirectID, err := rr.Save(ctx, &model.Redirect{
		Code: code,
		URL:  expectedURL,
	})
	redirect, err = rr.FindByCode(ctx, code)
	if redirect.Code != code {
		t.Errorf("RedirectRepo.FindByCode() return redirect with code %s , expected %s", redirect.Code, code)
	}
	if redirect.URL != expectedURL {
		t.Errorf("RedirectRepo.FindByCode() return redirect with URL %s , expected %s", redirect.URL, expectedURL)
	}

	redirect2, err := rr.FindByID(ctx, redirectID)
	if redirect2.URL != expectedURL {
		t.Errorf("RedirectRepo.FindByID() return redirect with URL %s , expected %s", redirect.URL, expectedURL)
	}
	expectedCount := int64(1)
	count, err := rr.Count(ctx)
	if count != expectedCount {
		t.Errorf("RedirectRepo.Count() return %d , expected %d", count, expectedCount)
	}
	expectedURL2 := "http://www.example.com"
	code2 := shortid.MustGenerate()
	_, err = rr.Save(ctx, &model.Redirect{
		Code: code2,
		URL:  expectedURL2,
	})
	redirects, err := rr.FindAll(ctx, &repository.RedirectListOptions{
		Query: "",
		LimitOffset: &repository.LimitOffset{
			Limit:  5,
			Offset: 0,
		},
	})
	expectedCount2 := 2
	if len(redirects) != expectedCount2 {
		t.Errorf("RedirectRepo.FindAll() return %d records , expected %d", len(redirects), expectedCount2)
	}
	err = rr.Delete(ctx, int64(1))
	if err != nil {
		t.Errorf("RedirectRepo.Delete() return no err, expected %s", err.Error())
	}
	expectedCount3 := int64(1)
	count3, err := rr.Count(ctx)
	if count3 != expectedCount3 {
		t.Errorf("RedirectRepo.Count() return %d , expected %d", count3, expectedCount3)
	}

	if err != nil {
		t.Errorf("error occurs: %s", err.Error())
	}

	util.MigrationDown(cfg, WD)
}
