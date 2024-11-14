package server

import (
	"github.com/RaceSimHub/race-hub-backend/internal/config"
	"github.com/RaceSimHub/race-hub-backend/internal/database"
	serverNotification "github.com/RaceSimHub/race-hub-backend/internal/server/routes/notification"
	serviceNotification "github.com/RaceSimHub/race-hub-backend/internal/service/notification"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/swag/example/basic/docs"

	swaggerFiles "github.com/swaggo/files"
)

// Server serves http requests for the service
type Server struct {
	Port   string
	Router *gin.Engine
}

// NewServer creates a new http server and setup routing
func NewServer() (s Server) {
	if config.ENVIRONMENT != "DEV" {
		gin.SetMode(gin.ReleaseMode)
	}

	s.Port = config.ServerPort
	s.Router = s.setupRouter()

	return
}

// Start runs the http server on a specific address
func (s Server) Start() {
	address := ":" + s.Port
	err := s.Router.Run(address)
	if err != nil {
		panic(err)
	}
}

func (Server) setupRouter() (router *gin.Engine) {
	router = gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders: []string{"Origin", "Content-Length", "Content-Type", "Accept", "Authorization"},
	}))

	docs.SwaggerInfo.BasePath = "/" + config.ApiVersion
	docs.SwaggerInfo.Host = config.SwaggerServerHost

	router.GET("/docs/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	routerGroup := router.Group(config.ApiVersion)

	notification := serverNotification.NewNotification(*serviceNotification.NewNotification(database.DbQuerier))
	routerGroup.POST("/notification", notification.Post)
	routerGroup.PUT("/notification/:id", notification.Put)
	routerGroup.DELETE("/notification/:id", notification.Delete)
	routerGroup.GET("/notification/last", notification.GetLastMessage)
	routerGroup.GET("/notification", notification.GetList)

	return
}
