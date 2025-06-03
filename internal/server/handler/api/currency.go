package api

import (
	"log/slog"
	"net/http"

	"github.com/UnLess24/coin/client/internal/dto"
	coinserver "github.com/UnLess24/coin/client/internal/server/coin_server"
	"github.com/gin-gonic/gin"
)

// Currency godoc
// @Summary	информация о конкретной валюте
// @Schemes
// @Description возвращает информацию о конкретной валюте
// @Tags Валюты
// @Param Authorization header string true "Authorization Bearer header"
// @securitydefinitions.oauth2.password OAuth2Password
// @in header
// @name Authorization
// @tokenUrl
// @Accept json
// @Produce json
// @Param currency query string false "наименование валюты"
// @Success 200 {array} coinserver.QuoteResponse
// @Failure 400 {object} dto.ResponseError
// @Failure 401 {object} middleware.AuthMiddlewareError
// @Router /api/currency [get]
func Currency(coinSrv coinserver.CoinServer) gin.HandlerFunc {
	return func(c *gin.Context) {
		currency := c.Query("currency")
		if currency == "" {
			slog.Error("currency is empty")
			c.JSON(http.StatusBadRequest, dto.ResponseError{ErrorMessage: "currency is empty"})
			return
		}

		data, err := coinSrv.Currency(c.Request.Context(), currency)
		if err != nil {
			slog.Error("failed to get currency", "error", err)
			c.JSON(http.StatusInternalServerError, dto.ResponseError{ErrorMessage: "failed to get currency"})
			return
		}

		c.JSON(http.StatusOK, data)
	}
}
