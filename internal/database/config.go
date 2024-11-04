package database

import (
	"database/sql"

	"github.com/RaceSimHub/race-hub-backend/internal/config"
	"github.com/RaceSimHub/race-hub-backend/internal/database/sqlc"
	"github.com/sirupsen/logrus"

	_ "github.com/lib/pq"
)

var (
	DbQuerier sqlc.Querier
)

type Config struct{}

func (c Config) Start() {
	conn := c.setup(
		config.DatabaseDriver,
		config.DatabaseUser,
		config.DatabasePass,
		config.DatabaseName,
		config.DatabasePort,
		config.DatabaseHost,
	)
	DbQuerier = sqlc.New(conn)
}

func (Config) setup(dbDriver, dbUser, dbPassword, dbName, dbPort, dbHost string) *sql.DB {
	dbSource := "host=" + dbHost + " port=" + dbPort + " user=" + dbUser + " password=" + dbPassword + " dbname=" + dbName + " sslmode=disable"

	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		logrus.Error(err.Error())
		panic(err.Error())
	}

	return conn
}
