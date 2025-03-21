package config

import "os"

var (
	Environment = os.Getenv("ENVIRONMENT")

	DatabaseDriver = os.Getenv("DATABASE_DRIVER")
	DatabaseUser   = os.Getenv("DATABASE_USER")
	DatabasePass   = os.Getenv("DATABASE_PASS")
	DatabaseName   = os.Getenv("DATABASE_NAME")
	DatabasePort   = os.Getenv("DATABASE_PORT")
	DatabaseHost   = os.Getenv("DATABASE_HOST")

	ServerPort = os.Getenv("SERVER_PORT")

	ApiVersion = os.Getenv("API_VERSION")

	SwaggerServerHost = os.Getenv("SWAGGER_SERVER_HOST")

	JwtSecret = os.Getenv("JWT_SECRET")

	IRacingEmail    = os.Getenv("IRACING_EMAIL")
	IRacingPassword = os.Getenv("IRACING_PASSWORD")

	EmailFrom     = os.Getenv("EMAIL_FROM")
	EmailPassword = os.Getenv("EMAIL_PASSWORD")
	EmailHost     = os.Getenv("EMAIL_HOST")
	EmailPort     = os.Getenv("EMAIL_PORT")

	RaceHubHost = os.Getenv("RACE_HUB_HOST")
)
