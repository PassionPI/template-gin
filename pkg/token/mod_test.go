package token_test

import (
	"testing"

	"app_ink/pkg/token"
)

const secret = "secret"
const fake = "fake"

var tokenInstance *token.Token = token.New(secret)
var username = "username"

func TestJWT(t *testing.T) {
	tokenStr, err := tokenInstance.Generate(username)
	if err != nil {
		t.Error(err)
	}
	claims, parsedToken, err := tokenInstance.Parse(tokenStr)
	if err != nil {
		t.Error(err)
	}
	if claims.Username != username {
		t.Error("username not match")
	}
	if !parsedToken.Valid {
		t.Error("token not valid")
	}
}
