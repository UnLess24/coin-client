package server

import (
	"net/http"

	"github.com/UnLess24/coin/client/config"
	"github.com/UnLess24/coin/client/internal/database"
	coinserver "github.com/UnLess24/coin/client/internal/server/coin_server"
	"github.com/UnLess24/coin/client/internal/server/handler"
	"github.com/UnLess24/coin/client/internal/server/handler/api"
	"github.com/UnLess24/coin/client/internal/server/middleware"
	"github.com/gin-gonic/gin"
)

type Server struct {
	*http.Server
	db      database.DB
	coinSrv coinserver.CoinServer
}

func New(addr string, db database.DB, coinSrv coinserver.CoinServer, cfg *config.Config) *Server {
	if cfg.Server.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()

	r.GET("/healthcheck", handler.HealthCheck)
	r.POST("/login", handler.Login(db, []byte(cfg.JWTSecretKey)))
	r.POST("/register", handler.Register(db))

	apiRoute := r.Group("/api")
	{
		apiRoute.Use(middleware.Auth([]byte(cfg.JWTSecretKey)))
		apiRoute.GET("/currencies", api.Currencies(coinSrv))
		apiRoute.GET("/currency", api.Currency(coinSrv))
	}

	return &Server{
		Server: &http.Server{
			Addr:    addr,
			Handler: r,
		},
		db:      db,
		coinSrv: coinSrv,
	}
}
