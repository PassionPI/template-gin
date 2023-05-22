package token

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Claims 结构体
// 包含了 JWT 的 payload
type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// Token 结构体
// 包含了 JWT 的配置 过期时间 & 刷新时间
type Token struct {
	secret  []byte
	Expire  time.Duration
	Expired time.Duration
	Refresh time.Duration
}

// New 创建一个 Token 实例
func New(secret string) *Token {
	return &Token{
		secret:  []byte(secret),
		Expire:  7 * 24 * time.Hour,
		Expired: 0 * time.Hour,
		Refresh: 24 * time.Hour,
	}
}

// Parse 解析 JWT token
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

// Generate 生成 JWT token
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
