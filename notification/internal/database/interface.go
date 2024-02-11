package database

import (
	"github.com/Impisigmatus/PestControlExpert/notification/autogen"
	"github.com/Impisigmatus/PestControlExpert/notification/internal/models"
	"github.com/jmoiron/sqlx"
)

//go:generate mockgen -source=interface.go -package mocks -destination ../../autogen/mocks/database.go
type Database interface {
	GetSubscribers() ([]models.Subscriber, error)
	AddSubscriber(subscriber models.Subscriber) (bool, error)
	PushNotification(tx *sqlx.Tx, notification autogen.Notification) error
	GetTX() (*sqlx.Tx, error)
}
