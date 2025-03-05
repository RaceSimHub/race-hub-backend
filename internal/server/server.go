package server

import (
	"github.com/RaceSimHub/race-hub-backend/internal/config"
	"github.com/RaceSimHub/race-hub-backend/internal/database"
	"github.com/RaceSimHub/race-hub-backend/internal/middleware"
	serverDriver "github.com/RaceSimHub/race-hub-backend/internal/server/routes/driver"
	serverLogin "github.com/RaceSimHub/race-hub-backend/internal/server/routes/login"
	serverNotification "github.com/RaceSimHub/race-hub-backend/internal/server/routes/notification"
	serverTemplate "github.com/RaceSimHub/race-hub-backend/internal/server/routes/template"
	serverTrack "github.com/RaceSimHub/race-hub-backend/internal/server/routes/track"
	serviceDriver "github.com/RaceSimHub/race-hub-backend/internal/service/driver"
	serviceNotification "github.com/RaceSimHub/race-hub-backend/internal/service/notification"
	serviceTrack "github.com/RaceSimHub/race-hub-backend/internal/service/track"
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
	if config.Environment != "DEV" {
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

	login := serverLogin.NewLogin(*serviceUser.NewUser(database.DbQuerier))
	router.POST("/login", login.PostLogin)
	router.POST("/logout", login.PostLogout)
	router.GET("/login", login.GetLogin)

	authRouterGroup := router.Use(middleware.JWTMiddleware())

	template := serverTemplate.NewTemplate(database.DbQuerier)
	authRouterGroup.GET("/", template.Home)

	driver := serverDriver.NewDriver(*serviceDriver.NewDriver(database.DbQuerier))
	authRouterGroup.GET("/drivers", driver.GetList)
	authRouterGroup.POST("/drivers", driver.Post)
	authRouterGroup.GET("/drivers/:id", driver.GetByID)
	authRouterGroup.PUT("/drivers/:id", driver.Put)
	authRouterGroup.GET("/drivers/new", driver.New)
	authRouterGroup.GET("/drivers/delete/:id", driver.Delete)

	track := serverTrack.NewTrack(*serviceTrack.NewTrack(database.DbQuerier))
	authRouterGroup.GET("/tracks", track.GetList)
	authRouterGroup.POST("/tracks", track.Post)
	authRouterGroup.GET("/tracks/:id", track.GetByID)
	authRouterGroup.PUT("/tracks/:id", track.Put)
	authRouterGroup.GET("/tracks/new", track.New)
	authRouterGroup.GET("/tracks/delete/:id", track.Delete)

	notification := serverNotification.NewNotification(*serviceNotification.NewNotification(database.DbQuerier))
	authRouterGroup.POST("/notifications", notification.Post)
	authRouterGroup.PUT("/notifications/:id", notification.Put)
	authRouterGroup.DELETE("/notifications/:id", notification.Delete)
	authRouterGroup.GET("/notifications/last", notification.GetLastMessage)
	authRouterGroup.GET("/notifications", notification.GetList)

	return
}
