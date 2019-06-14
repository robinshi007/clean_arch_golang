package web

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	validator "gopkg.in/go-playground/validator.v9"

	"github.com/robinshi007/goweb/db"
	"github.com/robinshi007/goweb/domain/model"
	"github.com/robinshi007/goweb/domain/repository"
	"github.com/robinshi007/goweb/interface/persistence/postgres"
	"github.com/robinshi007/goweb/util"
)

// NewRouter
func NewUserRouter(uHandler *UserHandler) http.Handler {
	r := chi.NewRouter()
	r.Get("/", uHandler.Fetch)
	r.Get("/{id:[0-9]+}", uHandler.GetById)
	r.Post("/", uHandler.Create)
	r.Put("/{id:[0-9]+}", uHandler.Update)
	r.Delete("/{id:[0-9]+}", uHandler.Delete)
	return r
}

// NewPostHandler ...
func NewUserHandler(db *db.Db) *UserHandler {
	return &UserHandler{
		repo: postgres.NewUserRepo(db),
	}
}

// Post ...
type UserHandler struct {
	repo repository.UserRepository
}

// Fetch all post data
func (u *UserHandler) Fetch(w http.ResponseWriter, r *http.Request) {
	res, _ := u.repo.Fetch(5)

	util.RespondWithJSON(w, http.StatusOK, res)
}

// Create a new post
func (u *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	user := model.User{}
	json.NewDecoder(r.Body).Decode(&user)

	if ok, err := isRequestValid(&user); !ok {
		util.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	newId, err := u.repo.Create(&user)
	fmt.Println(newId)
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
	data := model.User{Id: int64(id)}
	json.NewDecoder(r.Body).Decode(&data)
	res, err := u.repo.Update(&data)

	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, "Server Error")
	} else {
		util.RespondWithJSON(w, http.StatusOK, res)
	}

}

// GetById returns a post details
func (u *UserHandler) GetById(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	res, err := u.repo.GetById(int64(id))

	if err != nil {
		util.RespondWithError(w, http.StatusNoContent, "Content not found")
	} else {
		util.RespondWithJSON(w, http.StatusOK, res)
	}
}

// Delete a post
func (u *UserHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	res, err := u.repo.Delete(int64(id))

	if !res && err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, "Server Error")
	} else {
		util.RespondWithJSON(w, http.StatusMovedPermanently, map[string]string{"message": "Delete Successfully"})
	}

}

func isRequestValid(u *model.User) (bool, error) {
	validate := validator.New()
	err := validate.Struct(u)
	if err != nil {
		return false, err
	}
	return true, nil
}
