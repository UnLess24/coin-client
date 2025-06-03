package handler

import (
	"context"
	"net/http"
	"net/mail"
	"time"

	"github.com/UnLess24/coin/client/internal/database"
	"github.com/UnLess24/coin/client/internal/dto"
	"github.com/UnLess24/coin/client/internal/models/user"
	"github.com/gin-gonic/gin"
)

var (
	invalidRequest = "invalid request"
	invalidEmail   = "invalid email"
)

// Register godoc
// @Summary регистрация пользователя
// @Schemes
// @Description выполняет регистрацию
// @Tags Регистрация и авторизация
// @Param request body dto.RegisterRequest true "request"
// @Accept json
// @Produce json
// @Success 201
// @Failure 400 {object} dto.ResponseError
// @Router /register [post]
func Register(db database.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		req := dto.RegisterRequest{}
		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, dto.ResponseError{ErrorMessage: invalidRequest})
			return
		}

		if _, err := mail.ParseAddress(req.Email); err != nil {
			c.JSON(http.StatusBadRequest, dto.ResponseError{ErrorMessage: invalidRequest})
			return
		}

		ctx, cancel := context.WithTimeout(c.Request.Context(), time.Second*5)
		defer cancel()
		err := db.CreateUser(ctx, user.FromRegisterRequest(req))
		if err != nil {
			c.JSON(http.StatusBadRequest, dto.ResponseError{ErrorMessage: err.Error()})
			return
		}

		c.Status(http.StatusCreated)
	}
}
