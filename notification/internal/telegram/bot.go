package telegram

import (
	"fmt"

	"github.com/Impisigmatus/PestControlExpert/notification/internal/database"
	"github.com/Impisigmatus/PestControlExpert/notification/internal/models"
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
)

type Bot struct {
	api  *tg.BotAPI
	db   database.Database
	pass string
}

func NewBot(cfg database.PostgresConfig, token string, pass string) *Bot {
	api, err := tg.NewBotAPI(token)
	if err != nil {
		logrus.Panicf("Invalid telegram bot: %s", err)
	}

	bot := &Bot{
		api:  api,
		db:   database.NewPostgres(cfg),
		pass: pass,
	}

	bot.consume()
	return bot
}

func (bot *Bot) Notify(notification string) error {
	tx, err := bot.db.GetTX()
	if err != nil {
		return fmt.Errorf("Invalid transaction: %s", err)
	}

	if err := bot.db.PushNotification(tx, notification); err != nil {
		if txErr := tx.Rollback(); txErr != nil {
			err = fmt.Errorf("%s with invalid rollback: %s", err, txErr)
		}

		return fmt.Errorf("Invalid db push notification: %s", err)
	}

	if err := bot.send(notification); err != nil {
		if txErr := tx.Rollback(); txErr != nil {
			err = fmt.Errorf("%s with invalid rollback: %s", err, txErr)
		}

		return fmt.Errorf("Invalid tg bot send notification: %s", err)
	}

	if err := tx.Commit(); err != nil {
		if txErr := tx.Rollback(); txErr != nil {
			err = fmt.Errorf("%s with invalid rollback: %s", err, txErr)
		}

		return fmt.Errorf("Invalid commit: %s", err)
	}

	return nil
}

func (bot *Bot) send(msg string) error {
	subscribers, err := bot.db.GetSubscribers()
	if err != nil {
		return fmt.Errorf("Invalid subscribers: %s", err)
	}

	for _, subscriber := range subscribers {
		msg := tg.NewMessage(subscriber.ChatID, msg)
		if _, err := bot.api.Send(msg); err != nil {
			return fmt.Errorf("Invalid send msg for %s@%s[%d]: %s", subscriber.Name, subscriber.Username, subscriber.ChatID, err)
		}
	}

	return nil
}

func (bot *Bot) consume() {
	go func() {
		updater := tg.NewUpdate(0)
		updater.Timeout = 60
		updates := bot.api.GetUpdatesChan(updater)

		for update := range updates {
			if update.Message != nil {
				if update.Message.Text == bot.pass {
					ok, err := bot.db.AddSubscriber(models.Subscriber{
						ChatID:   update.Message.Chat.ID,
						Username: update.Message.From.UserName,
						Name:     update.Message.From.FirstName,
					})
					if err != nil {
						logrus.Errorf("Invalid add subscriber: %s", err)
						continue
					}

					var text string
					if ok {
						text = "Вы уже были подписаны на оповещения от сервисов сайта Pest Control Expert"
					} else {
						text = "Вы успешно подписаны на оповещения от сервисов сайта Pest Control Expert"
					}

					msg := tg.NewMessage(update.Message.Chat.ID, text)
					msg.ReplyToMessageID = update.Message.MessageID

					if _, err := bot.api.Send(msg); err != nil {
						logrus.Errorf("Invalid send msg: %s", err)
						continue
					}
				}
			}
		}
	}()
}
