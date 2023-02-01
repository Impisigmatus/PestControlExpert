package database

import (
	"github.com/Impisigmatus/PestControlExpert/notification/internal/models"
	"github.com/jmoiron/sqlx"
)

type Database interface {
	GetSubscribers() ([]models.Subscriber, error)
	AddSubscriber(subscriber models.Subscriber) (bool, error)
	GetTX() (*sqlx.Tx, error)
}
