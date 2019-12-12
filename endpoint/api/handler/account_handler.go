package handler

import (
	"context"
	"fmt"
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

// NewAccountRouter -
func NewAccountRouter(uHandler *AccountHandler) http.Handler {
	r := chi.NewRouter()
	r.Get("/", uHandler.GetAll)
	r.Get("/{id:[0-9]+}", uHandler.GetByID)
	r.Get("/{name:[a-zA-Z0-9@]+}/by_email", uHandler.GetByEmail)
	r.Post("/", uHandler.Create)
	r.Put("/{id:[0-9]+}", uHandler.UpdatePassword)
	r.Delete("/{id:[0-9]+}", uHandler.Delete)
	return r
}

// NewAccountHandler -
func NewAccountHandler() *AccountHandler {
	repo := postgres.NewAccountRepo()
	pre := presenter.NewAccountPresenter()
	return &AccountHandler{
		uc:  ctn.NewAccountUseCase(repo, pre, 2*time.Second),
		rsp: respond.NewRespond("json"),
	}
}

// AccountHandler -
type AccountHandler struct {
	uc  usecase.AccountUsecase
	rsp api.Respond
}

// GetAll the post data
func (u *AccountHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	res, err := u.uc.GetAll(context.Background(), 5)
	if err != nil {
		u.rsp.Error(w, err)
	} else {
		u.rsp.OK(w, res)
	}
}

// Create a new post
func (u *AccountHandler) Create(w http.ResponseWriter, r *http.Request) {
	account := in.NewAccount{}
	u.rsp.Decode(r.Body, &account)

	newID, err := u.uc.Create(context.Background(), &account)
	if err != nil {
		fmt.Println("create err:", err)
		u.rsp.Error(w, err)
	} else {
		u.rsp.Created(w, newID)
	}
}

// UpdatePassword - update a post by id
func (u *AccountHandler) UpdatePassword(w http.ResponseWriter, r *http.Request) {
	account := in.EditAccount{
		ID: chi.URLParam(r, "id"),
	}
	u.rsp.Decode(r.Body, &account)

	res, err := u.uc.UpdatePassword(context.Background(), &account)
	if err != nil {
		u.rsp.Error(w, err)
	} else {
		u.rsp.OK(w, res)
	}

}

// GetByID returns a post details
func (u *AccountHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	res, err := u.uc.GetByID(context.Background(), &in.FetchAccount{ID: id})
	if err != nil {
		u.rsp.Error(w, err)
	} else {
		u.rsp.OK(w, res)
	}
}

// GetByEmail - returns a post details
func (u *AccountHandler) GetByEmail(w http.ResponseWriter, r *http.Request) {
	email := chi.URLParam(r, "email")
	res, err := u.uc.GetByEmail(context.Background(), &in.FetchAccountByEmail{Email: email})
	if err != nil {
		u.rsp.Error(w, err)
	} else {
		u.rsp.OK(w, res)
	}
}

// Delete a post
func (u *AccountHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	err := u.uc.Delete(context.Background(), &in.FetchAccount{ID: id})
	if err != nil {
		u.rsp.Error(w, err)
	} else {
		u.rsp.OK(w, id)
	}
}
