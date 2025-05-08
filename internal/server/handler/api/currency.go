package api

import (
	"log/slog"
	"net/http"

	coinserver "github.com/UnLess24/coin/client/internal/server/coin_server"
	"github.com/gin-gonic/gin"
)

func Currency(coinSrv coinserver.CoinServer) gin.HandlerFunc {
	return func(c *gin.Context) {
		currency := c.Query("currency")
		if currency == "" {
			slog.Error("currency is empty")
			c.JSON(http.StatusBadRequest, gin.H{
				"errorMessage": "currency is empty",
			})
			return
		}

		data, err := coinSrv.Currency(c.Request.Context(), currency)
		if err != nil {
			slog.Error("failed to get currency", "error", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"errorMessage": "failed to get currency",
			})
			return
		}

		c.JSON(http.StatusOK, data)
	}
}
