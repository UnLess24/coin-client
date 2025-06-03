package api

import (
	"log/slog"
	"net/http"

	"github.com/UnLess24/coin/client/internal/dto"
	coinserver "github.com/UnLess24/coin/client/internal/server/coin_server"
	"github.com/gin-gonic/gin"
)

// Currencies godoc
// @Summary	список валют
// @Schemes
// @Description возвращает список валют
// @Tags Валюты
// @Param Authorization header string true "Authorization Bearer header"
// @securitydefinitions.oauth2.password OAuth2Password
// @in header
// @name Authorization
// @tokenUrl
// @Accept json
// @Produce json
// @Success 200 {array} coinserver.Currency
// @Failure 400 {object} dto.ResponseError
// @Failure 401 {object} middleware.AuthMiddlewareError
// @Router /api/currencies [get]
func Currencies(coinSrv coinserver.CoinServer) gin.HandlerFunc {
	return func(c *gin.Context) {
		data, err := coinSrv.List(c.Request.Context())
		if err != nil {
			slog.Error("failed to get currencies", "error", err)
			c.JSON(http.StatusInternalServerError, dto.ResponseError{ErrorMessage: "failed to get currencies"})
			return
		}

		c.JSON(http.StatusOK, data)
	}
}
