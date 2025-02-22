package handler

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/UnLess24/coin/client/internal/database"
	"github.com/UnLess24/coin/client/internal/dto"
	"github.com/gin-gonic/gin"

	"github.com/golang-jwt/jwt/v5"
)

var errBadCredentials = "user or password is incorrect"

func Login(db database.DB, JWTSecreteKey []byte) gin.HandlerFunc {
	return func(c *gin.Context) {
		req := dto.LoginRequest{}
		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"errorMessage": errBadCredentials,
			})
			return
		}

		u, err := db.FindUserByEmail(req.Email, req.Password)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"errorMessage": errBadCredentials,
			})
			return
		}

		// Генерируем JWT-токен для пользователя,
		// который он будет использовать в будущих HTTP-запросах

		// Генерируем полезные данные, которые будут храниться в токене
		payload := jwt.MapClaims{
			"sub": u.Email,
			"exp": time.Now().Add(time.Hour * 24).Unix(),
		}

		// Создаем новый JWT-токен и подписываем его по алгоритму HS256
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

		t, err := token.SignedString(JWTSecreteKey)
		if err != nil {
			slog.Error("JWT token signing", "errorMessage", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"errorMessage": errBadCredentials,
			})
			return
		}

		c.JSON(http.StatusOK, dto.LoginResponse{AccessToken: t})
	}
}
