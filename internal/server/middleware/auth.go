package middleware

import (
	"net/http"

	jwttoken "github.com/UnLess24/coin/client/internal/server/jwt_token"
	"github.com/gin-gonic/gin"
)

type AuthMiddlewareError struct {
	Error string `json:"error"`
}

func Auth(JWTSecretKey []byte) gin.HandlerFunc {
	return func(c *gin.Context) {
		bearerToken := c.GetHeader("Authorization")
		if bearerToken == "" {
			c.JSON(http.StatusUnauthorized, AuthMiddlewareError{Error: "authorization required"})
			c.Abort()
			return
		}

		err := jwttoken.Parse(bearerToken, JWTSecretKey)
		if err != nil {
			c.JSON(http.StatusUnauthorized, AuthMiddlewareError{Error: "invalid token"})
			c.Abort()
			return
		}

		c.Next()
	}
}
