package server

import (
	"net/http"

	"github.com/UnLess24/coin/client/internal/database"
	"github.com/UnLess24/coin/client/internal/server/handler"
	"github.com/gin-gonic/gin"
)

type Server struct {
	*http.Server
	DB database.DB
}

func New(addr string, db database.DB) *Server {
	r := gin.Default()

	r.GET("/healthcheck", handler.HealthCheck)

	return &Server{
		Server: &http.Server{
			Addr:    addr,
			Handler: r,
		},
		DB: db,
	}
}
