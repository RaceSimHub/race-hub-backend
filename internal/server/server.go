package server

import (
	"github.com/RaceSimHub/race-hub-backend/internal/config"
	"github.com/RaceSimHub/race-hub-backend/internal/database"
	"github.com/RaceSimHub/race-hub-backend/internal/middleware"
	serverNotification "github.com/RaceSimHub/race-hub-backend/internal/server/routes/notification"
	serverUser "github.com/RaceSimHub/race-hub-backend/internal/server/routes/user"
	serviceNotification "github.com/RaceSimHub/race-hub-backend/internal/service/notification"
	serviceUser "github.com/RaceSimHub/race-hub-backend/internal/service/user"
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

	freeRouterGroup := router.Group(config.ApiVersion)

	user := serverUser.NewUser(*serviceUser.NewUser(database.DbQuerier))
	freeRouterGroup.POST("/login", user.PostLogin)

	authRouterGroup := freeRouterGroup.Use(middleware.JWTMiddleware())
	authRouterGroup.POST("/users", user.Post)

	notification := serverNotification.NewNotification(*serviceNotification.NewNotification(database.DbQuerier))
	authRouterGroup.POST("/notifications", notification.Post)
	authRouterGroup.PUT("/notifications/:id", notification.Put)
	authRouterGroup.DELETE("/notifications/:id", notification.Delete)
	authRouterGroup.GET("/notifications/last", notification.GetLastMessage)
	authRouterGroup.GET("/notifications", notification.GetList)

	return
}
