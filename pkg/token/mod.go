package token

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

type Token struct {
	secret  []byte
	Expire  time.Duration
	Expired time.Duration
	Refresh time.Duration
}

func New(secret string) *Token {
	return &Token{
		secret:  []byte(secret),
		Expire:  7 * 24 * time.Hour,
		Expired: 0 * time.Hour,
		Refresh: 24 * time.Hour,
	}
}

func (auth *Token) Parse(token string) (*Claims, *jwt.Token, error) {
	claims := &Claims{}
	parsed, err := jwt.ParseWithClaims(
		token,
		claims,
		func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return auth.secret, nil
		},
	)

	return claims, parsed, err
}

func (auth *Token) Generate(username string) (string, error) {
	return jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		&Claims{
			Username: username,
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(auth.Expire)),
			},
		},
	).SignedString(auth.secret)
}
