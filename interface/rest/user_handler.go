package rest

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi"

	"clean_arch/infra/database"
	"clean_arch/interface/postgres"
	"clean_arch/pkg/util"
	"clean_arch/presenter"
	"clean_arch/usecase"
	"clean_arch/usecase/in"
)

// NewUserRouter -
func NewUserRouter(uHandler *UserHandler) http.Handler {
	r := chi.NewRouter()
	r.Get("/", uHandler.Fetch)
	r.Get("/{id:[0-9]+}", uHandler.GetByID)
	r.Get("/{name:[a-z0-9]+}/by_name", uHandler.GetByName)
	r.Post("/", uHandler.Create)
	r.Put("/{id:[0-9]+}", uHandler.Update)
	r.Delete("/{id:[0-9]+}", uHandler.Delete)
	return r
}

// NewUserHandler -
func NewUserHandler(dbm database.DB) *UserHandler {
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
func (u *UserHandler) Fetch(w http.ResponseWriter, r *http.Request) {
	res, _ := u.uc.Fetch(context.Background(), 5)

	util.RespondWithJSON(w, http.StatusOK, res)
}

// Create a new post
func (u *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	user := in.PostUser{}
	json.NewDecoder(r.Body).Decode(&user)

	if err := user.Validate(); err != nil {
		util.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	newID, err := u.uc.Create(context.Background(), &user)
	fmt.Println(newID)
	if err != nil {
		//respondWithError(w, http.StatusInternalServerError, "Server Error")
		util.RespondWithError(w, util.GetStatusCode(err), err.Error())
		fmt.Println(err.Error())
	} else {
		util.RespondWithJSON(w, http.StatusCreated, map[string]string{"message": "Successfully Created"})
	}
}

// Update a post by id
func (u *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	user := in.PutUser{ID: int64(id)}
	json.NewDecoder(r.Body).Decode(&user)

	if err := user.Validate(); err != nil {
		util.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	// check exist by ID
	_, err := u.uc.GetByID(context.Background(), user.ID)
	if err != nil {
		util.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	res, err := u.uc.Update(context.Background(), &user)
	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, err.Error())
	} else {
		util.RespondWithJSON(w, http.StatusOK, res)
	}

}

// GetByID returns a post details
func (u *UserHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	res, err := u.uc.GetByID(context.Background(), int64(id))

	if err != nil {
		util.RespondWithError(w, http.StatusNoContent, "Content not found")
	} else {
		util.RespondWithJSON(w, http.StatusOK, res)
	}
}

// GetByName returns a post details
func (u *UserHandler) GetByName(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	res, err := u.uc.GetByName(context.Background(), name)

	if err != nil {
		util.RespondWithError(w, http.StatusNoContent, "Content not found")
	} else {
		util.RespondWithJSON(w, http.StatusOK, res)
	}
}

// Delete a post
func (u *UserHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	err := u.uc.Delete(context.Background(), int64(id))

	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, err.Error())
	} else {
		util.RespondWithJSON(w, http.StatusMovedPermanently, map[string]string{"message": "Delete Successfully"})
	}

}
