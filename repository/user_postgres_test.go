package repository_test

import (
	"context"
	"testing"

	"github.com/robinshi007/goweb/db"
	"github.com/robinshi007/goweb/model"
	"github.com/robinshi007/goweb/repository"
)

func TestUserCRUD(t *testing.T) {
	conn, err := db.NewDb("localhost", "5432", "postgres", "postgres", "test")
	if err != nil {
		panic(err)
	}
	ctx := context.WithValue(context.Background(), "hi", "Golang")
	userRepo := repository.NewUserRepo(conn)
	id, err := userRepo.Create(ctx, &model.User{Name: "Rob", Desc: "Pike"})
	if err != nil {
		panic(err)
	}
	user, err := userRepo.GetById(ctx, id)
	if err != nil {
		panic(err)
	}
	if user.Id != id {
		t.Errorf("id is incorrect, got: %d, want: %d", user.Id, id)
	}

	id2, err := userRepo.Create(ctx, &model.User{Name: "Hello", Desc: "World"})
	if err != nil {
		panic(err)
	}
	users, err := userRepo.Fetch(ctx, int64(5))
	if len(users) != 2 {
		t.Errorf("Fetch() results size is incorrect, got: %d, want: %d", len(users), 2)
	}

	user.Name = "Robin"
	_, err = userRepo.Update(ctx, user)
	if err != nil {
		panic(err)
	}
	user2, err := userRepo.GetById(ctx, user.Id)
	if user2.Name != "Robin" {
		t.Errorf("Update() u.name is incorrect, got: %v, want: %v", user2.Name, "Robin")
	}
	for _, rowId := range []int64{id, id2} {
		_, err := userRepo.Delete(ctx, rowId)
		if err != nil {
			panic(err)
		}
	}
	users2, err := userRepo.Fetch(ctx, int64(5))
	if len(users2) != 0 {
		t.Errorf("Fetch() results size is incorrect, got: %d, want: %d", len(users2), 0)
	}
}
