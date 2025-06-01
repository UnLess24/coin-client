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
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Server struct {
	*http.Server
	db      database.DB
	coinSrv coinserver.CoinServer
}

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8081
// @BasePath  /

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
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
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return &Server{
		Server: &http.Server{
			Addr:    addr,
			Handler: r,
		},
		db:      db,
		coinSrv: coinSrv,
	}
}
