package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/robinshi007/goweb/db"
	"github.com/robinshi007/goweb/model"
	"github.com/robinshi007/goweb/repository"
)

// NewPostHandler ...
func NewUserHandler(db *db.Db) *User {
	return &User{
		repo: repository.NewUserRepo(db),
	}
}

// Post ...
type User struct {
	repo repository.UserRepo
}

// Fetch all post data
func (u *User) Fetch(w http.ResponseWriter, r *http.Request) {
	res, _ := u.repo.Fetch(r.Context(), 5)

	respondwithJSON(w, http.StatusOK, res)
}

// Create a new post
func (u *User) Create(w http.ResponseWriter, r *http.Request) {
	user := model.User{}
	json.NewDecoder(r.Body).Decode(&user)

	newId, err := u.repo.Create(r.Context(), &user)
	fmt.Println(newId)
	if err != nil {
		//respondWithError(w, http.StatusInternalServerError, "Server Error")
		respondWithError(w, http.StatusInternalServerError, err.Error())
	} else {
		respondwithJSON(w, http.StatusCreated, map[string]string{"message": "Successfully Created"})
	}
}

// Update a post by id
func (u *User) Update(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	data := model.User{Id: int64(id)}
	json.NewDecoder(r.Body).Decode(&data)
	res, err := u.repo.Update(r.Context(), &data)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Server Error")
	}

	respondwithJSON(w, http.StatusOK, res)
}

// GetById returns a post details
func (u *User) GetById(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	res, err := u.repo.GetById(r.Context(), int64(id))

	if err != nil {
		respondWithError(w, http.StatusNoContent, "Content not found")
	}

	respondwithJSON(w, http.StatusOK, res)
}

// Delete a post
func (u *User) Delete(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	_, err := u.repo.Delete(r.Context(), int64(id))

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Server Error")
	}

	respondwithJSON(w, http.StatusMovedPermanently, map[string]string{"message": "Delete Successfully"})
}

// respondwithJSON write json response format
func respondwithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// respondwithError return error message
func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondwithJSON(w, code, map[string]string{"message": msg})
}
