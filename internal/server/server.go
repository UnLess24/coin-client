package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func New(addr string) *http.Server {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	return &http.Server{
		Addr:    addr,
		Handler: r,
	}
}
