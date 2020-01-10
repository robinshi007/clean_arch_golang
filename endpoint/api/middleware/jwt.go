package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"

	"clean_arch/domain/model"
	"clean_arch/endpoint/api/globals"
)

// AccountClaims -
type AccountClaims struct {
	Email string `json:"email"`
	Name  string `json:"name"`
	jwt.StandardClaims
}

var (
	jwtKey = []byte("my_secret_key")
)

// contextKey is a value for use with context.WithValue. It's used as
// a pointer so it fits in an interface{} without allocation.
type contextKey struct {
	email string
}

func (k *contextKey) String() string {
	return "jwt auth context value " + k.email
}

// Context keys
var (
	TokenCtxKey = &contextKey{"Token"}
	ErrorCtxKey = &contextKey{"Error"}
)

// TokenInfo -
type TokenInfo struct {
	Name      string `json:"name"`
	Token     string `json:"token"`
	ExpiresAt int64  `json:"expiresAt"`
}

// GenerateToken -
func GenerateToken(email string, name string) (TokenInfo, error) {
	expiresAt := time.Now().Add(30 * time.Minute)
	issuedBy := "nobody"

	// define payload
	claim := AccountClaims{
		email,
		name,
		jwt.StandardClaims{
			ExpiresAt: expiresAt.Unix(),
			Issuer:    issuedBy,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	ss, err := token.SignedString([]byte(jwtKey))
	return TokenInfo{name, ss, expiresAt.Unix()}, err

}

// ParseToken -
func ParseToken(ss string) (*jwt.Token, *AccountClaims, error) {
	if ss == "" {
		return nil, nil, model.ErrTokenEmpty
	}

	// parse payload
	token, err := jwt.ParseWithClaims(ss, &AccountClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtKey), nil
	})

	// validate payload
	if err == nil && token != nil {
		if claim, ok := token.Claims.(*AccountClaims); ok && token.Valid {
			return token, claim, nil
		}
		return nil, nil, model.ErrTokenInvalid
	}

	v, _ := err.(*jwt.ValidationError)
	if v.Errors == jwt.ValidationErrorExpired {
		return nil, nil, model.ErrTokenExpired
	}
	return nil, nil, model.ErrTokenInvalid
}

// NewJWTContext -
func NewJWTContext(ctx context.Context, t *jwt.Token, err error) context.Context {
	ctx = context.WithValue(ctx, TokenCtxKey, t)
	ctx = context.WithValue(ctx, ErrorCtxKey, err)
	return ctx
}

// FromJWTContext -
func FromJWTContext(ctx context.Context) (*jwt.Token, *AccountClaims, error) {
	token, _ := ctx.Value(TokenCtxKey).(*jwt.Token)
	var claims *AccountClaims
	if token != nil {
		if tokenClaims, ok := token.Claims.(*AccountClaims); ok {
			claims = tokenClaims
		} else {
			return nil, nil, fmt.Errorf("jwt: unknown type of Claims: %s", token.Claims)
		}
	} else {
		claims = &AccountClaims{}
	}
	err, _ := ctx.Value(ErrorCtxKey).(error)
	return token, claims, err
}

// TokenFromHTTPRequest -
func TokenFromHTTPRequest(r *http.Request) string {
	reqToken := r.Header.Get("Authorization")
	var tokenString string
	splitToken := strings.Split(reqToken, "Bearer")
	if len(splitToken) == 2 {
		tokenString = strings.TrimSpace(splitToken[1])
	}
	return tokenString
}

// WithJWTVerify - check token and put claims into context
func JWTVerify() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		hfn := func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()

			//token, err := VerifyRequest(ja, r, findTokenFns...)
			tokenString := TokenFromHTTPRequest(r)
			token, _, err := ParseToken(tokenString)
			newCtx := NewJWTContext(ctx, token, err)
			next.ServeHTTP(w, r.WithContext(newCtx))
		}
		return http.HandlerFunc(hfn)
	}
}

// JWTAuthenticator - a authentication middleware to enforce access from the
// Verifier middleware request context values. The Authenticator sends a 401 Unauthorized
// response for any unverified tokens and passes the good ones through.
func JWTAuthenticator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _, err := FromJWTContext(r.Context())

		if err != nil {
			globals.Respond.Error(w, err)
			return
		}

		// load user info according to claims info

		// Token is authenticated, pass it through
		next.ServeHTTP(w, r)
	})
}
