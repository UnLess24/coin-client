package api

import (
	"log/slog"
	"net/http"

	coinserver "github.com/UnLess24/coin/client/internal/server/coin_server"
	"github.com/gin-gonic/gin"
)

func Currencies(coinSrv coinserver.CoinServer) gin.HandlerFunc {
	return func(c *gin.Context) {
		data, err := coinSrv.List(c.Request.Context())
		if err != nil {
			slog.Error("failed to get currencies", "error", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"errorMessage": "failed to get currencies",
			})
			return
		}

		c.JSON(http.StatusOK, data)
	}
}
