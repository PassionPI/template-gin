package middleware

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

const jwtBearer = "Bearer "
const jwtBearerLen = len(jwtBearer)

// AuthValidator validates the JWT token
func (mid *Middleware) AuthValidator() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorization := ctx.Request.Header.Get("Authorization")

		if authorization == "" {
			ctx.AbortWithStatusJSON(
				http.StatusForbidden,
				gin.H{
					"error": "Missing authorization header",
				},
			)
			return
		}

		if len(authorization) < jwtBearerLen || authorization[:jwtBearerLen] != jwtBearer {
			ctx.AbortWithStatusJSON(
				http.StatusBadRequest,
				gin.H{
					"error": "Invalid authorization header format",
				},
			)
			return
		}

		// Parse the JWT token and verify the signature
		claims, token, err := mid.core.Token.Parse(authorization[jwtBearerLen:])

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				ctx.AbortWithStatusJSON(
					http.StatusUnauthorized,
					gin.H{
						"error": "Invalid token signature",
					},
				)
				return
			}
			ctx.AbortWithStatusJSON(
				http.StatusBadRequest,
				gin.H{"error": err.Error()},
			)
			return
		}

		if !token.Valid {
			ctx.AbortWithStatusJSON(
				http.StatusUnauthorized,
				gin.H{
					"error": "Invalid token",
				},
			)
			return
		}

		if claims.ExpiresAt.Time.Sub(time.Now()) < mid.core.Token.Refresh {
			newToken, _ := mid.core.Token.Generate(claims.Username)
			ctx.Header("Authorization", newToken)
		}

		// Store the user ID in the context for later use
		ctx.Set("username", claims.Username)

		ctx.Next()
	}
}
