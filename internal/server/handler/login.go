package handler

import (
	"context"
	"log/slog"
	"net/http"
	"time"

	"github.com/UnLess24/coin/client/internal/database"
	"github.com/UnLess24/coin/client/internal/dto"
	"github.com/gin-gonic/gin"

	"github.com/golang-jwt/jwt/v5"
)

var errBadCredentials = "user or password is incorrect"

// Login godoc
// @Summary авторизация пользователя
// @Schemes
// @Description выполняет авторизацию
// @Tags Регистрация и авторизация
// @Param request body dto.LoginRequest true "request"
// @Accept json
// @Produce json
// @Success 200 {object} dto.LoginResponse
// @Failure 400 {object} dto.ResponseError
// @Router /login [post]
func Login(db database.DB, JWTSecreteKey []byte) gin.HandlerFunc {
	return func(c *gin.Context) {
		req := dto.LoginRequest{}
		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, dto.ResponseError{ErrorMessage: errBadCredentials})
			return
		}

		ctx, cancel := context.WithTimeout(c.Request.Context(), time.Second*5)
		defer cancel()
		u, err := db.FindUserByEmail(ctx, req.Email, req.Password)
		if err != nil {
			c.JSON(http.StatusBadRequest, dto.ResponseError{ErrorMessage: err.Error()})
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
		if err != nil || len(JWTSecreteKey) == 0 {
			slog.Error("JWT token signing", "errorMessage", err, "JWTSecreteKeyEmpty", len(JWTSecreteKey) == 0)
			c.JSON(http.StatusInternalServerError, dto.ResponseError{ErrorMessage: errBadCredentials})
			return
		}

		c.JSON(http.StatusOK, dto.LoginResponse{AccessToken: t})
	}
}
