package handler

import (
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

// NewUserRouter -
func NewUserRouter(uHandler *UserHandler) http.Handler {
	r := chi.NewRouter()
	r.Get("/", uHandler.FindAll)
	r.Get("/{id:[0-9]+}", uHandler.FindByID)
	r.Get("/{name:[a-zA-Z0-9]+}/by_name", uHandler.FindByName)
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
		uc:  ctn.NewUserUsecase(repo, pre),
		rsp: respond.NewRespond(registry.Cfg.Serializer.Code),
	}
}

// UserHandler -
type UserHandler struct {
	uc  usecase.UserUsecase
	rsp api.Responder
}

// FindAll the post data
func (u *UserHandler) FindAll(w http.ResponseWriter, r *http.Request) {
	res, err := u.uc.FindAll(r.Context(), &in.FetchAllOptions{})
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

	newID, err := u.uc.Create(r.Context(), &user)
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

	res, err := u.uc.Update(r.Context(), &user)
	if err != nil {
		u.rsp.Error(w, err)
	} else {
		u.rsp.OK(w, res)
	}

}

// FindByID returns a post details
func (u *UserHandler) FindByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	res, err := u.uc.FindByID(r.Context(), &in.FetchUser{ID: id})
	if err != nil {
		u.rsp.Error(w, err)
	} else {
		u.rsp.OK(w, res)
	}
}

// FindByName returns a post details
func (u *UserHandler) FindByName(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	res, err := u.uc.FindByName(r.Context(), &in.FetchUserByName{Name: name})
	if err != nil {
		u.rsp.Error(w, err)
	} else {
		u.rsp.OK(w, res)
	}
}

// Delete a post
func (u *UserHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	err := u.uc.Delete(r.Context(), &in.FetchUser{ID: id})
	if err != nil {
		u.rsp.Error(w, err)
	} else {
		u.rsp.OK(w, id)
	}
}
