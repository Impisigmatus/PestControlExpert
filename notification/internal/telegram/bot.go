package telegram

import (
	"fmt"

	"github.com/Impisigmatus/PestControlExpert/notification/internal/database"
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
)

type Bot struct {
	api  *tg.BotAPI
	db   database.Database
	pass string
}

func NewBot(token string, pass string) *Bot {
	api, err := tg.NewBotAPI(token)
	if err != nil {
		logrus.Panicf("Invalid telegram bot: %s", err)
	}

	bot := &Bot{
		api:  api,
		db:   database.NewMemDB(),
		pass: pass,
	}

	bot.consume()
	return bot
}

func (bot *Bot) Send(msg string) error {
	ids, err := bot.db.GetSubscribers()
	if err != nil {
		return fmt.Errorf("Invalid subscribers: %s", err)
	}

	for _, id := range ids {
		msg := tg.NewMessage(id, msg)
		if _, err := bot.api.Send(msg); err != nil {
			return fmt.Errorf("Invalid send msg: %s", err)
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
					if err := bot.db.AddSubscriber(update.Message.Chat.ID); err != nil {
						logrus.Panicf("Invalid add subscriber: %s", err)
					}

					msg := tg.NewMessage(update.Message.Chat.ID, "Вы успешно подписаны на оповещения от сервисов сайта Pest Control Expert")
					msg.ReplyToMessageID = update.Message.MessageID
					if _, err := bot.api.Send(msg); err != nil {
						logrus.Panicf("Invalid send msg: %s", err)
					}
				}
			}
		}
	}()
}
