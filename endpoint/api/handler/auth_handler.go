package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/go-chi/chi"

	"clean_arch/adapter/postgres"
	"clean_arch/adapter/presenter"
	"clean_arch/domain/model"
	"clean_arch/domain/usecase"
	"clean_arch/domain/usecase/in"
	"clean_arch/endpoint/api"
	"clean_arch/endpoint/api/middleware"
	"clean_arch/endpoint/api/respond"
	ctn "clean_arch/usecase"
)

// NewAuthRouter -
func NewAuthRouter(authHandler *AuthHandler) http.Handler {
	r := chi.NewRouter()
	r.Post("/login", authHandler.Login)
	r.Get("/refresh_token", authHandler.RefreshToken)
	return r
}

// NewAuthHandler -
func NewAuthHandler() *AuthHandler {
	repo := postgres.NewAccountRepo()
	pre := presenter.NewAccountPresenter()
	return &AuthHandler{
		uc:  ctn.NewAccountUseCase(repo, pre, 2*time.Second),
		rsp: respond.NewRespond("json"),
	}
}

// AuthHandler -
type AuthHandler struct {
	uc  usecase.AccountUsecase
	rsp api.Respond
}

// Login -
func (a *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	// clean the cookie
	account := in.LoginAccountByEmail{}
	a.rsp.Decode(r.Body, &account)
	res, err := a.uc.Login(context.Background(), &account)
	if err != nil {
		a.rsp.Error(w, err)
	} else if !res {
		a.rsp.Error(w, model.ErrAccountNotMatch)
	} else {
		tokenString, err := middleware.GenerateToken(account.Email)
		if err != nil {
			a.rsp.Error(w, err)
		}
		http.SetCookie(w, &http.Cookie{
			Name:    "token",
			Value:   tokenString,
			Expires: time.Now().Add(30 * time.Minute),
		})
		a.rsp.OK(w, res)
	}
}

// RefreshToken -
func (a *AuthHandler) RefreshToken(w http.ResponseWriter, r *http.Request) {

	token, _, err := middleware.FromContext(r.Context())

	if err != nil {
		http.Error(w, http.StatusText(401), 401)
		return
	}

	if token == nil || !token.Valid {
		http.Error(w, http.StatusText(401), 401)
		return
	}
	if claim, ok := token.Claims.(*middleware.AccountClaims); ok && token.Valid {
		if (claim.ExpiresAt - time.Now().Unix()) < 60*5 {
			// generate new jwt token and set cookie
			a.rsp.OK(w, "refresh token")
			return
		}
		a.rsp.OK(w, "no need to refresh token")
		return
	}
	http.Error(w, http.StatusText(401), 401)
	return
}
