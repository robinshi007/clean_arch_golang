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
		uc:  ctn.NewAccountUseCase(repo, pre, 2*time.Second),
		rsp: respond.NewRespond(registry.Cfg.Serializer.Code),
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
		a.rsp.Error(w, model.ErrAuthNotMatch)
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
		a.rsp.Error(w, model.ErrTokenIsInvalid)
		return
	}

	if token == nil || !token.Valid {
		a.rsp.Error(w, model.ErrTokenIsInvalid)
		return
	}
	if claim, ok := token.Claims.(*middleware.AccountClaims); ok && token.Valid {
		if (claim.ExpiresAt - time.Now().Unix()) < 60*10 {
			// generate new jwt token and set cookie
			tokenString, err := middleware.GenerateToken(claim.Email)
			if err != nil {
				a.rsp.Error(w, err)
				return
			}
			http.SetCookie(w, &http.Cookie{
				Name:    "token",
				Value:   tokenString,
				Expires: time.Now().Add(30 * time.Minute),
			})
			a.rsp.OK(w, "token is refreshed")
			return
		}
		a.rsp.OK(w, "token is no need to refresh")
		return
	}
	a.rsp.Error(w, model.ErrTokenIsInvalid)
	return
}

// JWTAuthenticator - a authentication middleware to enforce access from the
// Verifier middleware request context values. The Authenticator sends a 401 Unauthorized
// response for any unverified tokens and passes the good ones through.
func (a *AuthHandler) JWTAuthenticator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, _, err := middleware.FromContext(r.Context())
		if err != nil || token == nil || !token.Valid {
			a.rsp.Error(w, model.ErrTokenIsInvalid)
			return
		}

		// Token is authenticated, pass it through
		next.ServeHTTP(w, r)
	})
}
