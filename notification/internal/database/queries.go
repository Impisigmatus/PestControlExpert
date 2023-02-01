package database

import (
	"fmt"

	"github.com/Impisigmatus/PestControlExpert/notification/autogen"
	"github.com/Impisigmatus/PestControlExpert/notification/internal/models"
	"github.com/jmoiron/sqlx"
)

func (pg *Postgres) GetSubscribers() ([]models.Subscriber, error) {
	const query = "SELECT chat_id, username, name FROM main.subscribers;"

	subscribers := make([]models.Subscriber, 0)
	if err := pg.db.Select(&subscribers, query); err != nil {
		return nil, fmt.Errorf("Invalid SELECT main.subscribers: %s", err)
	}

	return subscribers, nil
}

func (pg *Postgres) AddSubscriber(subscriber models.Subscriber) (bool, error) {
	const query = `
INSERT INTO main.subscribers (
	chat_id,
	username,
	name
) VALUES (
	:chat_id,
	:username,
	:name
) ON CONFLICT (chat_id) DO NOTHING;`

	result, err := pg.db.NamedExec(query, subscriber)
	if err != nil {
		return false, fmt.Errorf("Invalid INSERT main.subscribers: %s", err)
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return false, fmt.Errorf("Invalid rows affected: %s", err)
	}

	return affected == 0, nil
}

func (pg *Postgres) PushNotification(tx *sqlx.Tx, notification autogen.Notification) error {
	return nil
}
