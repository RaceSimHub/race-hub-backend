package server

import (
	"log"
	"os"
	"path/filepath"

	"github.com/RaceSimHub/race-hub-backend/internal/config"
	"github.com/RaceSimHub/race-hub-backend/internal/database"
	"github.com/RaceSimHub/race-hub-backend/internal/middleware"
	serverDriver "github.com/RaceSimHub/race-hub-backend/internal/server/routes/driver"
	serverDriverLink "github.com/RaceSimHub/race-hub-backend/internal/server/routes/driver_link"
	serverNotification "github.com/RaceSimHub/race-hub-backend/internal/server/routes/notification"
	serverTemplate "github.com/RaceSimHub/race-hub-backend/internal/server/routes/template"
	serverTrack "github.com/RaceSimHub/race-hub-backend/internal/server/routes/track"
	serverUser "github.com/RaceSimHub/race-hub-backend/internal/server/routes/user"
	serviceDriver "github.com/RaceSimHub/race-hub-backend/internal/service/driver"
	serviceDriverLink "github.com/RaceSimHub/race-hub-backend/internal/service/driver_link"
	serviceNotification "github.com/RaceSimHub/race-hub-backend/internal/service/notification"
	serviceTrack "github.com/RaceSimHub/race-hub-backend/internal/service/track"
	serviceUser "github.com/RaceSimHub/race-hub-backend/internal/service/user"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/swag/example/basic/docs"

	swaggerFiles "github.com/swaggo/files"
)

type Server struct {
	Port   string
	Router *gin.Engine
}

func NewServer() (s Server) {
	if config.Environment != "DEV" {
		gin.SetMode(gin.ReleaseMode)
	}

	s.Port = config.ServerPort
	s.Router = s.setupRouter()

	return
}

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

	basePath, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	staticPath := filepath.Join(basePath, "..", "static")

	router.Static("/js", filepath.Join(staticPath, "js"))
	router.Static("/css", filepath.Join(staticPath, "css"))
	router.StaticFile("/favicon.ico", filepath.Join(staticPath, "favicon.ico"))

	user := serverUser.NewUser(*serviceUser.NewUser(database.DbQuerier))
	router.GET("/login", user.GetLogin)
	router.POST("/login", user.PostLogin)
	router.POST("/logout", user.PostLogout)
	router.GET("/sign-up", user.GetSignUp)
	router.POST("/sign-up", user.PostUser)
	router.GET("/email-confirm", user.GetEmailConfirm)
	router.POST("/verify-code", user.PostEmailVerify)
	router.GET("/forgot-password", user.GetForgotPassword)
	router.POST("/forgot-password", user.PostForgotPassword)
	router.GET("/reset-password", user.GetResetPassword)
	router.POST("/reset-password", user.PostResetPassword)
	router.POST("/resend-email-confirmation", user.PostResendEmailConfirmation)

	authRouterGroup := router.Use(middleware.JWTMiddleware())

	template := serverTemplate.NewTemplate(database.DbQuerier)
	authRouterGroup.GET("/", template.Home)

	serviceDriver := serviceDriver.NewDriver(database.DbQuerier)

	driverLink := serverDriverLink.NewDriverLink(*serviceDriverLink.NewDriverLink(database.DbQuerier), *serviceDriver)
	authRouterGroup.GET("/driver/link", driverLink.GetDriverLink)
	authRouterGroup.GET("/driver/link/pending", driverLink.Pending)
	authRouterGroup.POST("/driver/link/:driverID", driverLink.Post)

	authAdminRouterGroup := authRouterGroup.Use(middleware.AdminMiddleware())

	driver := serverDriver.NewDriver(*serviceDriver)
	authAdminRouterGroup.GET("/admin/drivers", driver.GetList)
	authAdminRouterGroup.POST("/admin/drivers", driver.Post)
	authAdminRouterGroup.GET("/admin/drivers/:id", driver.GetByID)
	authAdminRouterGroup.PUT("/admin/drivers/:id", driver.Put)
	authAdminRouterGroup.GET("/admin/drivers/new", driver.New)
	authAdminRouterGroup.DELETE("/admin/drivers/:id", driver.Delete)
	authAdminRouterGroup.PUT("/admin/drivers/:id/irating", driver.UpdateIrating)

	track := serverTrack.NewTrack(*serviceTrack.NewTrack(database.DbQuerier))
	authAdminRouterGroup.GET("/admin/tracks", track.GetList)
	authAdminRouterGroup.POST("/admin/tracks", track.Post)
	authAdminRouterGroup.GET("/admin/tracks/:id", track.GetByID)
	authAdminRouterGroup.PUT("/admin/tracks/:id", track.Put)
	authAdminRouterGroup.GET("/admin/tracks/new", track.New)
	authAdminRouterGroup.DELETE("/admin/tracks/:id", track.Delete)

	notification := serverNotification.NewNotification(*serviceNotification.NewNotification(database.DbQuerier))
	authAdminRouterGroup.POST("/notifications", notification.Post)
	authAdminRouterGroup.PUT("/notifications/:id", notification.Put)
	authAdminRouterGroup.DELETE("/notifications/:id", notification.Delete)
	authAdminRouterGroup.GET("/notifications/last", notification.GetLastMessage)
	authAdminRouterGroup.GET("/notifications", notification.GetList)

	return
}
