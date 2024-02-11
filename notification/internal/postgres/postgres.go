package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Impisigmatus/PestControlExpert/notification/autogen"
	"github.com/Impisigmatus/PestControlExpert/notification/internal/models"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

//go:generate mockgen -source=postgres.go -package mocks -destination ../../autogen/mocks/postgres.go
type IDatabase interface {
	GetSubscribers() ([]models.Subscriber, error)
	AddSubscriber(subscriber models.Subscriber) (bool, error)
	PushNotification(tx *sqlx.Tx, notification autogen.Notification) error
	GetTX() (*sqlx.Tx, error)
	CommitTX(tx *sqlx.Tx) error
	RollbackTX(tx *sqlx.Tx) error
}

type PostgresConfig struct {
	Hostname string
	Port     uint64
	Database string
	User     string
	Password string
}

type Postgres struct {
	db *sqlx.DB
}

const driver = "pgx"

func NewPostgres(cfg PostgresConfig) IDatabase {
	return &Postgres{db: sqlx.NewDb(newPostgres(cfg), driver)}
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
	db, err := sql.Open(driver, connection)
	if err != nil {
		logrus.Panicf("Invalid postgres connect: %s", err)
	}

	if err := db.Ping(); err != nil {
		logrus.Panicf("Invalid postgres ping: %s", err)
	}

	return db
}

func (pg *Postgres) GetTX() (*sqlx.Tx, error) {
	return pg.db.BeginTxx(context.Background(), nil)
}

func (pg *Postgres) CommitTX(tx *sqlx.Tx) error {
	return tx.Commit()
}

func (pg *Postgres) RollbackTX(tx *sqlx.Tx) error {
	return tx.Rollback()
}
