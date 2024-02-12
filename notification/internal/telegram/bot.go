package telegram

import (
	"fmt"

	"github.com/Impisigmatus/PestControlExpert/notification/autogen"
	"github.com/Impisigmatus/PestControlExpert/notification/internal/models"
	"github.com/Impisigmatus/PestControlExpert/notification/internal/postgres"
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
)

//go:generate mockgen -source=bot.go -package mocks -destination ../../autogen/mocks/telegram.go
type ITelegramAPI interface {
	Send(c tg.Chattable) (tg.Message, error)
	GetUpdatesChan(cfg tg.UpdateConfig) tg.UpdatesChannel
}

type Bot struct {
	db   postgres.IDatabase
	api  ITelegramAPI
	pass string
}

func NewBot(cfg postgres.Config, token string, pass string) *Bot {
	api, err := tg.NewBotAPI(token)
	if err != nil {
		logrus.Panicf("Invalid telegram bot: %s", err)
	}

	bot := newBot(postgres.NewPostgres(cfg), api, pass)
	bot.consume()
	return bot
}

func newBot(db postgres.IDatabase, api ITelegramAPI, pass string) *Bot {
	return &Bot{
		db:   db,
		api:  api,
		pass: pass,
	}
}

func (bot *Bot) Notify(notification autogen.Notification) error {
	tx, err := bot.db.GetTX()
	if err != nil {
		return fmt.Errorf("Invalid transaction: %s", err)
	}

	if err := bot.db.PushNotification(tx, notification); err != nil {
		if txErr := bot.db.RollbackTX(tx); txErr != nil {
			err = fmt.Errorf("%s with invalid rollback: %s", err, txErr)
		}

		return fmt.Errorf("Invalid db push notification: %s", err)
	}

	description := ""
	if notification.Description != nil {
		description = *notification.Description
	}
	msg := fmt.Sprintf(
		"ФИО: %s\nТелефон: %s\nКоментарий: %s",
		notification.Name,
		notification.Phone,
		description,
	)
	if err := bot.send(msg); err != nil {
		if txErr := bot.db.RollbackTX(tx); txErr != nil {
			err = fmt.Errorf("%s with invalid rollback: %s", err, txErr)
		}

		return fmt.Errorf("Invalid tg bot send notification: %s", err)
	}

	if err := bot.db.CommitTX(tx); err != nil {
		if txErr := bot.db.RollbackTX(tx); txErr != nil {
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
			if err := bot.handle(update); err != nil {
				logrus.Errorf("Invalid update: %s", err)
				continue
			}
		}
	}()
}

func (bot *Bot) handle(update tg.Update) error {
	if update.Message != nil {
		if update.Message.Text == bot.pass {
			ok, err := bot.db.AddSubscriber(models.Subscriber{
				ChatID:   update.Message.Chat.ID,
				Username: update.Message.From.UserName,
				Name:     update.Message.From.FirstName,
			})
			if err != nil {
				return fmt.Errorf("Invalid add subscriber: %s", err)
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
				return fmt.Errorf("Invalid send msg: %s", err)
			}
		}
	}
	return nil
}
