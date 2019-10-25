package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi"

	"clean_arch/adapter/postgres"
	"clean_arch/adapter/presenter"
	"clean_arch/domain/model"
	"clean_arch/infra"
	"clean_arch/interface/rest"
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

// GetAll the post data
func (u *UserHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	res, _ := u.uc.GetAll(context.Background(), 5)
	rest.RespondOK(w, res)
}

// Create a new post
func (u *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	user := input.PostUser{}
	json.NewDecoder(r.Body).Decode(&user)

	if err := user.Validate(); err != nil {
		rest.RespondError(w, model.ErrEntityBadInput)
		return
	}
	newID, err := u.uc.Create(context.Background(), &user)
	if err != nil {
		rest.RespondError(w, err)
	} else {
		rest.RespondCreated(w, newID)
	}
}

// Update a post by id
func (u *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	// check exist by ID
	user2, err := u.uc.GetByID(context.Background(), int64(id))
	if err != nil {
		rest.RespondError(w, model.ErrEntityBadInput)
		return
	}
	user := input.PutUser{
		User: &model.User{ID: int64(id), Name: user2.Name},
	}
	json.NewDecoder(r.Body).Decode(&user)

	if err := user.Validate(); err != nil {
		rest.RespondError(w, model.ErrEntityBadInput)
		return
	}

	res, err := u.uc.Update(context.Background(), &user)
	if err != nil {
		rest.RespondError(w, err)
	} else {
		rest.RespondOK(w, res)
	}

}

// GetByID returns a post details
func (u *UserHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	res, err := u.uc.GetByID(context.Background(), int64(id))

	if err != nil {
		rest.RespondError(w, err)
	} else {
		rest.RespondOK(w, res)
	}
}

// GetByName returns a post details
func (u *UserHandler) GetByName(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	res, err := u.uc.GetByName(context.Background(), name)

	if err != nil {
		rest.RespondError(w, err)
	} else {
		rest.RespondOK(w, res)
	}
}

// Delete a post
func (u *UserHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		rest.RespondError(w, model.ErrEntityBadInput)
	}
	err = u.uc.Delete(context.Background(), int64(id))
	if err != nil {
		rest.RespondError(w, err)
	} else {
		rest.RespondOK(w, string(id))
	}

}
