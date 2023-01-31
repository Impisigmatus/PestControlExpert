package database

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type PostgresConfig struct {
	Hostname string
	Port     uint32
	Database string
	User     string
	Password string
}

type Postgres struct {
	db *sqlx.DB
}

func NewPostgres(cfg PostgresConfig) Database {
	const driver = "pgx"
	return &Postgres{db: sqlx.NewDb(newPostgres(cfg), driver)}
}

func (pg *Postgres) GetSubscribers() ([]int64, error) {
	return nil, nil // TODO
}

func (pg *Postgres) AddSubscriber(id int64) (bool, error) {
	return false, nil // TODO
}

func newPostgres(cfg PostgresConfig) *sql.DB {
	pattern := fmt.Sprintf(
		"host=%s port=%d database=%s user=%s password=%s sslmode=disable",
		cfg.Hostname, cfg.Port, cfg.Database, cfg.User, cfg.Password,
	)

	config, err := pgx.ParseConfig(pattern)
	if err != nil {
		logrus.Panicf("Invalid postgres config: %s", err)
	}
	config.Logger = &pgxLogger{}

	connection := stdlib.RegisterConnConfig(config)
	db, err := sql.Open("pgx", connection)
	if err != nil {
		logrus.Panicf("Invalid postgres connect: %s", err)
	}

	if err := db.Ping(); err != nil {
		logrus.Panicf("Invalid postgres ping: %s", err)
	}

	return db
}

type pgxLogger struct{}

func (*pgxLogger) Log(ctx context.Context, level pgx.LogLevel, msg string, data map[string]interface{}) {
	var logger logrus.FieldLogger
	if data != nil {
		logger = logrus.StandardLogger().WithFields(data)
	} else {
		logger = logrus.StandardLogger()
	}

	switch level {
	case pgx.LogLevelTrace:
		logger.WithField("PGX_LOG_LEVEL", level).Debug(msg)
	case pgx.LogLevelDebug:
		logger.Debug(msg)
	case pgx.LogLevelInfo:
		logger.Info(msg)
	case pgx.LogLevelWarn:
		logger.Warn(msg)
	case pgx.LogLevelError:
		logger.Error(msg)
	default:
		logger.WithField("INVALID_PGX_LOG_LEVEL", level).Error(msg)
	}
}
