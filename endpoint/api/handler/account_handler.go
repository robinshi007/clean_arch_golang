package handler

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"

	"clean_arch/adapter/postgres"
	"clean_arch/adapter/presenter"
	"clean_arch/domain/usecase"
	"clean_arch/domain/usecase/in"
	"clean_arch/endpoint/api"
	"clean_arch/endpoint/api/respond"
	"clean_arch/registry"
	ctn "clean_arch/usecase"
)

// NewAccountRouter -
func NewAccountRouter(uHandler *AccountHandler) http.Handler {
	r := chi.NewRouter()
	r.Get("/", uHandler.FindAll)
	r.Get("/{id:[0-9]+}", uHandler.FindByID)
	r.Get("/{name:[a-zA-Z0-9@]+}/by_email", uHandler.FindByEmail)
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
		uc:  ctn.NewAccountUsecase(repo, pre),
		rsp: respond.NewRespond(registry.Cfg.Serializer.Code),
	}
}

// AccountHandler -
type AccountHandler struct {
	uc  usecase.AccountUsecase
	rsp api.Responder
}

// FindAll the post data
func (u *AccountHandler) FindAll(w http.ResponseWriter, r *http.Request) {
	res, err := u.uc.FindAll(r.Context(), &in.FetchAllOptions{})
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

	newID, err := u.uc.Create(r.Context(), &account)
	if err != nil {
		fmt.Println("create err:", err)
		u.rsp.Error(w, err)
	} else {
		u.rsp.Created(w, newID)
	}
}

// UpdatePassword - update a post by id
func (u *AccountHandler) UpdatePassword(w http.ResponseWriter, r *http.Request) {
	account := in.EditAccountPassword{
		ID: chi.URLParam(r, "id"),
	}
	u.rsp.Decode(r.Body, &account)

	res, err := u.uc.UpdatePassword(r.Context(), &account)
	if err != nil {
		u.rsp.Error(w, err)
	} else {
		u.rsp.OK(w, res)
	}

}

// FindByID returns a post details
func (u *AccountHandler) FindByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	res, err := u.uc.FindByID(r.Context(), &in.FetchAccount{ID: id})
	if err != nil {
		u.rsp.Error(w, err)
	} else {
		u.rsp.OK(w, res)
	}
}

// FindByEmail - returns a post details
func (u *AccountHandler) FindByEmail(w http.ResponseWriter, r *http.Request) {
	email := chi.URLParam(r, "email")
	res, err := u.uc.FindByEmail(r.Context(), &in.FetchAccountByEmail{Email: email})
	if err != nil {
		u.rsp.Error(w, err)
	} else {
		u.rsp.OK(w, res)
	}
}

// Delete a post
func (u *AccountHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	err := u.uc.Delete(r.Context(), &in.FetchAccount{ID: id})
	if err != nil {
		u.rsp.Error(w, err)
	} else {
		u.rsp.OK(w, id)
	}
}
