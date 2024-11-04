package main

import (
	"github.com/RaceSimHub/race-hub-backend/internal/database"
	"github.com/RaceSimHub/race-hub-backend/internal/server"
	"github.com/sirupsen/logrus"
)

func init() {
	configureLog()

	database.Config{}.Start()
}

func main() {
	server.NewServer().Start()
}

func configureLog() {
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:            true,
		DisableLevelTruncation: true,
		FullTimestamp:          true,
		TimestampFormat:        "2006/01/02 15:04:05",
	})
}
