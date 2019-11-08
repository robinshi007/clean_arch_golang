package handler

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi"

	"clean_arch/adapter/postgres"
	"clean_arch/adapter/presenter"
	"clean_arch/domain/model"
	"clean_arch/domain/usecase"
	"clean_arch/domain/usecase/in"
	"clean_arch/infra"
	"clean_arch/interface/api"
	"clean_arch/interface/api/respond"
	ctn "clean_arch/usecase"
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
	res, _ := u.uc.GetAll(context.Background(), 5)
	u.rsp.OK(w, res)
}

// Create a new post
func (u *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	user := in.PostUser{}
	//json.NewDecoder(r.Body).Decode(&user)
	u.rsp.Decode(r.Body, &user)

	if err := user.Validate(); err != nil {
		u.rsp.Error(w, model.ErrEntityBadInput)
		return
	}
	newID, err := u.uc.Create(context.Background(), &user)
	if err != nil {
		u.rsp.Error(w, err)
	} else {
		u.rsp.Created(w, newID)
	}
}

// Update a post by id
func (u *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	// check exist by ID
	user2, err := u.uc.GetByID(context.Background(), int64(id))
	if err != nil {
		u.rsp.Error(w, model.ErrEntityBadInput)
		return
	}
	user := in.PutUser{
		User: &model.User{ID: int64(id), Name: user2.Name},
	}
	u.rsp.Decode(r.Body, &user)

	if err := user.Validate(); err != nil {
		u.rsp.Error(w, model.ErrEntityBadInput)
		return
	}

	res, err := u.uc.Update(context.Background(), &user)
	if err != nil {
		u.rsp.Error(w, err)
	} else {
		u.rsp.OK(w, res)
	}

}

// GetByID returns a post details
func (u *UserHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	res, err := u.uc.GetByID(context.Background(), int64(id))

	if err != nil {
		u.rsp.Error(w, err)
	} else {
		u.rsp.OK(w, res)
	}
}

// GetByName returns a post details
func (u *UserHandler) GetByName(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	res, err := u.uc.GetByName(context.Background(), name)

	if err != nil {
		u.rsp.Error(w, err)
	} else {
		u.rsp.OK(w, res)
	}
}

// Delete a post
func (u *UserHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		u.rsp.Error(w, model.ErrEntityBadInput)
	}
	err = u.uc.Delete(context.Background(), int64(id))
	if err != nil {
		u.rsp.Error(w, err)
	} else {
		u.rsp.OK(w, string(id))
	}
}
