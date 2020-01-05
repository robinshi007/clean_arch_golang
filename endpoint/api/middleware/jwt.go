package middleware

import (
	"clean_arch/domain/model"
	"context"
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

// TokenInfo -
type TokenInfo struct {
	Token     string `json:"token"`
	ExpiresAt int64  `json:"expiresAt"`
}

// GenerateToken -
func GenerateToken(email string) (TokenInfo, error) {
	expiresAt := time.Now().Add(30 * time.Minute)
	issuedBy := "nobody"

	// define payload
	claim := AccountClaims{
		email,
		jwt.StandardClaims{
			ExpiresAt: expiresAt.Unix(),
			Issuer:    issuedBy,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	ss, err := token.SignedString([]byte(jwtKey))
	return TokenInfo{ss, expiresAt.Unix()}, err

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
