package server

import (
	"net/http"

	"github.com/UnLess24/coin/client/config"
	"github.com/UnLess24/coin/client/internal/database"
	"github.com/UnLess24/coin/client/internal/server/handler"
	"github.com/UnLess24/coin/client/internal/server/middleware"
	"github.com/gin-gonic/gin"
)

type Server struct {
	*http.Server
	DB database.DB
}

func New(addr string, db database.DB, cfg *config.Config) *Server {
	if cfg.Server.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()

	r.GET("/healthcheck", handler.HealthCheck)
	r.POST("/login", handler.Login(db, []byte(cfg.JWTSecretKey)))
	r.POST("/register", handler.Register(db))

	api := r.Group("/api")
	{
		api.Use(middleware.Auth([]byte(cfg.JWTSecretKey)))
	}

	return &Server{
		Server: &http.Server{
			Addr:    addr,
			Handler: r,
		},
		DB: db,
	}
}
