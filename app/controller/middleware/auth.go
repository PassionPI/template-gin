package middleware

import (
	"net/http"

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
				http.StatusUnauthorized,
				"Missing authorization header",
			)
			return
		}

		if len(authorization) < jwtBearerLen || authorization[:jwtBearerLen] != jwtBearer {
			ctx.AbortWithStatusJSON(
				http.StatusUnauthorized,
				"Invalid authorization header format",
			)
			return
		}

		// Parse the JWT token and verify the signature
		claims, token, err := mid.core.Dep.Token.Parse(authorization[jwtBearerLen:])

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				ctx.AbortWithStatusJSON(
					http.StatusUnauthorized,
					"Invalid token signature",
				)
				return
			}
			ctx.AbortWithStatusJSON(
				http.StatusUnauthorized,
				err.Error(),
			)
			return
		}

		if !token.Valid {
			ctx.AbortWithStatusJSON(
				http.StatusUnauthorized,
				"Invalid token",
			)
			return
		}

		// _, err = mid.core.Dep.Pg.UserFindByUsername(ctx, claims.Username)

		// if err != nil {
		// 	ctx.AbortWithStatusJSON(
		// 		http.StatusUnauthorized,
		// 		"Invalid token, user not found",
		// 	)
		// 	return
		// }

		// if claims.ExpiresAt.Time.Sub(time.Now()) < mid.core.Dep.Token.Refresh {
		// 	newToken, _ := mid.core.Dep.Token.Generate(claims.Username)
		// 	ctx.Header("Authorization", newToken)
		// }

		// Store the user ID in the context for later use
		ctx.Set("username", claims.Username)

		ctx.Next()
	}
}
