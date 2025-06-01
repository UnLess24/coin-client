package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HealthCheck godoc
// @Summary проверка работоспособности сервиса
// @Schemes
// @Description do check service
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} OK
// @Router /healthcheck [get]
func HealthCheck(c *gin.Context) {
	c.String(http.StatusOK, "OK")
}
