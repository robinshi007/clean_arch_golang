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
	rsp api.Respond
}

// Login -
func (a *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	// clean the cookie
	account := in.LoginAccountByEmail{}
	a.rsp.Decode(r.Body, &account)
	isMatched, err := a.uc.Login(r.Context(), &account)
	if err != nil {
		a.rsp.Error(w, err)
	} else if !isMatched {
		a.rsp.Error(w, model.ErrAuthNotMatch)
	} else {
		tokenInfo, err := middleware.GenerateToken(account.Email)
		if err != nil {
			a.rsp.Error(w, err)
		}
		a.rsp.OK(w, tokenInfo)
	}
}

// RefreshToken -
func (a *AuthHandler) RefreshToken(w http.ResponseWriter, r *http.Request) {

	_, claim, err := middleware.FromContext(r.Context())
	if err != nil {
		a.rsp.Error(w, err)
		return
	}

	// refresh time to 20S
	if (claim.ExpiresAt - time.Now().Unix()) < 5*60 {
		// generate new jwt token and set cookie
		tokenInfo, err := middleware.GenerateToken(claim.Email)
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

// JWTVerify - check token and put claims into context
func (a *AuthHandler) JWTVerify() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		hfn := func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()

			//token, err := VerifyRequest(ja, r, findTokenFns...)
			tokenString := middleware.TokenFromHTTPRequest(r)
			token, _, err := middleware.ParseToken(tokenString)
			newCtx := middleware.NewContext(ctx, token, err)
			next.ServeHTTP(w, r.WithContext(newCtx))
		}
		return http.HandlerFunc(hfn)
	}
}

// JWTAuthenticator - a authentication middleware to enforce access from the
// Verifier middleware request context values. The Authenticator sends a 401 Unauthorized
// response for any unverified tokens and passes the good ones through.
func (a *AuthHandler) JWTAuthenticator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _, err := middleware.FromContext(r.Context())

		if err != nil {
			a.rsp.Error(w, err)
			return
		}

		// load user info according to claims info

		// Token is authenticated, pass it through
		next.ServeHTTP(w, r)
	})
}
