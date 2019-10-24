package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi"

	"clean_arch/adapter/postgres"
	"clean_arch/adapter/presenter"
	"clean_arch/domain/model"
	"clean_arch/infra"
	"clean_arch/interface/rest/types"
	"clean_arch/usecase"
	"clean_arch/usecase/input"
)

// NewUserRouter -
func NewUserRouter(uHandler *UserHandler) http.Handler {
	r := chi.NewRouter()
	r.Get("/", uHandler.GetAll)
	r.Get("/{id:[0-9]+}", uHandler.GetByID)
	r.Get("/{name:[a-z0-9]+}/by_name", uHandler.GetByName)
	r.Post("/", uHandler.Create)
	r.Put("/{id:[0-9]+}", uHandler.Update)
	r.Delete("/{id:[0-9]+}", uHandler.Delete)
	return r
}

// NewUserHandler -
func NewUserHandler(dbm infra.DB) *UserHandler {
	repo := postgres.NewUserRepo(dbm)
	pre := presenter.NewUserPresenter()
	return &UserHandler{
		uc: usecase.NewUserUseCase(repo, pre, time.Second),
	}
}

// UserHandler -
type UserHandler struct {
	uc usecase.UserUsecase
}

// Fetch all post data
func (u *UserHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	res, _ := u.uc.GetAll(context.Background(), 5)

	types.RespondWithJSON(w, http.StatusOK, res)
}

// Create a new post
func (u *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	user := input.PostUser{}
	json.NewDecoder(r.Body).Decode(&user)

	if err := user.Validate(); err != nil {
		types.RespondWithError(w, "1101", err.Error())
		return
	}
	newID, err := u.uc.Create(context.Background(), &user)
	fmt.Println(newID)
	if err != nil {
		//respondWithError(w, http.StatusInternalServerError, "Server Error")
		types.RespondWithError(w, "1103", err.Error())
		fmt.Println(err.Error())
	} else {
		types.RespondWithJSON(w, http.StatusCreated, map[string]string{"message": "Successfully Created"})
	}
}

// Update a post by id
func (u *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	// check exist by ID
	user2, err := u.uc.GetByID(context.Background(), int64(id))
	if err != nil {
		types.RespondWithError(w, "1101", err.Error())
		return
	}
	user := input.PutUser{
		User: &model.User{ID: int64(id), Name: user2.Name},
	}
	json.NewDecoder(r.Body).Decode(&user)

	if err := user.Validate(); err != nil {
		types.RespondWithError(w, "1101", err.Error())
		return
	}

	res, err := u.uc.Update(context.Background(), &user)
	if err != nil {
		types.RespondWithError(w, "1103", err.Error())
	} else {
		types.RespondWithJSON(w, http.StatusOK, res)
	}

}

// GetByID returns a post details
func (u *UserHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	res, err := u.uc.GetByID(context.Background(), int64(id))

	if err != nil {
		types.RespondWithError(w, "1102", err.Error())
	} else {
		types.RespondWithJSON(w, http.StatusOK, res)
	}
}

// GetByName returns a post details
func (u *UserHandler) GetByName(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	res, err := u.uc.GetByName(context.Background(), name)

	if err != nil {
		types.RespondWithError(w, "1102", err.Error())
	} else {
		types.RespondWithJSON(w, http.StatusOK, res)
	}
}

// Delete a post
func (u *UserHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	err := u.uc.Delete(context.Background(), int64(id))

	if err != nil {
		types.RespondWithError(w, "1103", err.Error())
	} else {
		types.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "Delete Successfully"})
	}

}
