package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/go-chi/chi"

	"clean_arch/adapter/postgres"
	"clean_arch/adapter/presenter"
	"clean_arch/domain/usecase"
	"clean_arch/domain/usecase/in"
	"clean_arch/endpoint/api"
	"clean_arch/endpoint/api/respond"
	ctn "clean_arch/usecase"
)

// NewUserRouter -
func NewUserRouter(uHandler *UserHandler) http.Handler {
	r := chi.NewRouter()
	r.Get("/", uHandler.GetAll)
	r.Get("/{id:[0-9]+}", uHandler.GetByID)
	r.Get("/{name:[a-zA-Z0-9]+}/by_name", uHandler.GetByName)
	r.Post("/", uHandler.Create)
	r.Put("/{id:[0-9]+}", uHandler.Update)
	r.Delete("/{id:[0-9]+}", uHandler.Delete)
	return r
}

// NewUserHandler -
func NewUserHandler() *UserHandler {
	repo := postgres.NewUserRepo()
	pre := presenter.NewUserPresenter()
	return &UserHandler{
		uc:  ctn.NewUserUseCase(repo, pre, time.Second),
		rsp: respond.NewRespond("json"),
	}
}

// UserHandler -
type UserHandler struct {
	uc  usecase.UserUsecase
	rsp api.Respond
}

// GetAll the post data
func (u *UserHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	res, err := u.uc.GetAll(context.Background(), 5)
	if err != nil {
		u.rsp.Error(w, err)
	} else {
		u.rsp.OK(w, res)
	}
}

// Create a new post
func (u *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	user := in.NewUser{}
	u.rsp.Decode(r.Body, &user)

	newID, err := u.uc.Create(context.Background(), &user)
	if err != nil {
		u.rsp.Error(w, err)
	} else {
		u.rsp.Created(w, newID)
	}
}

// Update a post by id
func (u *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	user := in.EditUser{
		ID: chi.URLParam(r, "id"),
	}
	u.rsp.Decode(r.Body, &user)

	res, err := u.uc.Update(context.Background(), &user)
	if err != nil {
		u.rsp.Error(w, err)
	} else {
		u.rsp.OK(w, res)
	}

}

// GetByID returns a post details
func (u *UserHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	res, err := u.uc.GetByID(context.Background(), &in.FetchUser{ID: id})
	if err != nil {
		u.rsp.Error(w, err)
	} else {
		u.rsp.OK(w, res)
	}
}

// GetByName returns a post details
func (u *UserHandler) GetByName(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	res, err := u.uc.GetByName(context.Background(), &in.FetchUserByName{Name: name})
	if err != nil {
		u.rsp.Error(w, err)
	} else {
		u.rsp.OK(w, res)
	}
}

// Delete a post
func (u *UserHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	err := u.uc.Delete(context.Background(), &in.FetchUser{ID: id})
	if err != nil {
		u.rsp.Error(w, err)
	} else {
		u.rsp.OK(w, id)
	}
}
