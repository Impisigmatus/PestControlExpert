package database

import "github.com/Impisigmatus/PestControlExpert/notification/internal/models"

type Database interface {
	GetSubscribers() ([]models.Subscriber, error)
	AddSubscriber(subscriber models.Subscriber) (bool, error)
}
