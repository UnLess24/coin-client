package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HealthCheck godoc
// @Summary проверка работоспособности сервиса
// @Schemes
// @Description do check service
// @Tags HealthCheck
// @Accept plain
// @Produce plain
// @Success 200 string OK
// @Failure 400
// @Failure 500
// @Router /healthcheck [get]
func HealthCheck(c *gin.Context) {
	c.String(http.StatusOK, "OK")
}
