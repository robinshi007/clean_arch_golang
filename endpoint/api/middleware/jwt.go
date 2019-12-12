package middleware

import (
	"clean_arch/domain/model"
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// AccountClaims -
type AccountClaims struct {
	Email string `json:"email"`
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

// GenerateToken -
func GenerateToken(email string) (string, error) {
	expireAt := time.Now().Add(60 * time.Minute)
	issuedBy := "nobody"

	// define payload
	claim := AccountClaims{
		email,
		jwt.StandardClaims{
			ExpiresAt: expireAt.Unix(),
			Issuer:    issuedBy,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	ss, err := token.SignedString([]byte(jwtKey))
	return ss, err

}

// ParseToken -
func ParseToken(ss string) (*jwt.Token, *AccountClaims, error) {

	// parse payload
	token, err := jwt.ParseWithClaims(ss, &AccountClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtKey), nil
	})

	// validate payload
	if err == nil && token != nil {
		if claim, ok := token.Claims.(*AccountClaims); ok && token.Valid {
			return token, claim, nil
		}
		return nil, nil, model.ErrTokenIsInvalid
	}
	return nil, nil, model.ErrTokenIsInvalid
}

// NewContext -
func NewContext(ctx context.Context, t *jwt.Token, err error) context.Context {
	ctx = context.WithValue(ctx, TokenCtxKey, t)
	ctx = context.WithValue(ctx, ErrorCtxKey, err)
	return ctx
}

// FromContext -
func FromContext(ctx context.Context) (*jwt.Token, *AccountClaims, error) {
	token, _ := ctx.Value(TokenCtxKey).(*jwt.Token)
	var claims *AccountClaims
	if token != nil {
		if tokenClaims, ok := token.Claims.(*AccountClaims); ok {
			claims = tokenClaims
		} else {
			return nil, nil, errors.New(fmt.Sprintf("jwt: unknown type of Claims: %T", token.Claims))
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
	splitToken := strings.Split(reqToken, "Bearer ")
	if len(splitToken) > 1 {
		tokenString = splitToken[1]
	}
	//fmt.Println("tokenString", tokenString)
	return tokenString
}

// JWTVerify - check token and put claims into context
func JWTVerify() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		hfn := func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()

			//token, err := VerifyRequest(ja, r, findTokenFns...)
			tokenString := TokenFromHTTPRequest(r)
			token, _, err := ParseToken(tokenString)
			newCtx := NewContext(ctx, token, err)
			next.ServeHTTP(w, r.WithContext(newCtx))
		}
		return http.HandlerFunc(hfn)
	}
}
