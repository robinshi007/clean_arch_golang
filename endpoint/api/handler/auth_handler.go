package handler

import (
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
	"clean_arch/registry"
	ctn "clean_arch/usecase"
)

// NewAuthRouter -
func NewAuthRouter(authHandler *AuthHandler) http.Handler {
	r := chi.NewRouter()
	//r.Post("/login", authHandler.Login)
	r.Get("/refresh_token", authHandler.RefreshToken)
	return r
}

// NewAuthHandler -
func NewAuthHandler() *AuthHandler {
	repo := postgres.NewAccountRepo()
	pre := presenter.NewAccountPresenter()
	return &AuthHandler{
		uc:  ctn.NewAccountUseCase(repo, pre, 5*time.Second),
		rsp: respond.NewRespond(registry.Cfg.Serializer.Code),
	}
}

// AuthHandler -
type AuthHandler struct {
	uc  usecase.AccountUsecase
	rsp api.Responder
}

// Login -
func (a *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	// clean the cookie
	account := in.LoginAccountByEmail{}
	a.rsp.Decode(r.Body, &account)
	isMatched, name, err := a.uc.Login(r.Context(), &account)
	if err != nil {
		a.rsp.Error(w, err)
	} else if !isMatched {
		a.rsp.Error(w, model.ErrAuthNotMatch)
	} else {
		tokenInfo, err := middleware.GenerateToken(account.Email, name)
		if err != nil {
			a.rsp.Error(w, err)
		}
		a.rsp.OK(w, tokenInfo)
	}
}

// RefreshToken -
func (a *AuthHandler) RefreshToken(w http.ResponseWriter, r *http.Request) {

	_, claim, err := middleware.FromJWTContext(r.Context())
	if err != nil {
		a.rsp.Error(w, err)
		return
	}

	// refresh time to 20S
	if (claim.ExpiresAt - time.Now().Unix()) < 5*60 {
		// generate new jwt token and set cookie
		tokenInfo, err := middleware.GenerateToken(claim.Email, claim.Name)
		if err != nil {
			a.rsp.Error(w, err)
			return
		}
		a.rsp.OK(w, tokenInfo)
		return
	}
	a.rsp.OK(w, "token is no need to refresh")
	return
}
