package handler

import (
	"fmt"
	"net/http"
	"net/mail"

	"github.com/UnLess24/coin/client/internal/database"
	"github.com/UnLess24/coin/client/internal/dto"
	"github.com/UnLess24/coin/client/internal/models/user"
	"github.com/gin-gonic/gin"
)

var (
	invalidRequest = "invalid request"
	invalidEmail   = "invalid email"
)

func Register(db database.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		req := dto.RegisterRequest{}
		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"errorMessage": fmt.Sprintf("%s", invalidRequest),
			})
			return
		}

		if _, err := mail.ParseAddress(req.Email); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"errorMessage": fmt.Sprintf("%s", invalidEmail),
			})
			return
		}

		err := db.CreateUser(user.FromRegisterRequest(req))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"errorMessage": err.Error(),
			})
			return
		}

		c.Status(http.StatusCreated)
	}
}
