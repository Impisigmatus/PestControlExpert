package telegram

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/Impisigmatus/PestControlExpert/notification/autogen"
	"github.com/Impisigmatus/PestControlExpert/notification/autogen/mocks"
	"github.com/Impisigmatus/PestControlExpert/notification/internal/models"
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/golang/mock/gomock"
	"github.com/jmoiron/sqlx"
)

func Test_TelegramBot(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	db := mocks.NewMockIDatabase(ctrl)
	api := mocks.NewMockITelegramAPI(ctrl)
	bot := newBot(db, api, "test")

	t.Run("Notify()", func(t *testing.T) {
		sample_tx := &sqlx.Tx{Tx: &sql.Tx{}}
		sample_notification := autogen.Notification{
			Phone:       "+79991234567",
			Name:        "Василий",
			Description: nil,
		}
		t.Run("With subscribers", func(t *testing.T) {
			db.EXPECT().GetTX().Return(sample_tx, nil)
			db.EXPECT().PushNotification(sample_tx, sample_notification).Return(nil)
			db.EXPECT().CommitTX(sample_tx).Return(nil)
			db.EXPECT().GetSubscribers().Return([]models.Subscriber{{}}, nil)
			api.EXPECT().Send(gomock.Any()).Return(tg.Message{}, nil)

			if err := bot.Notify(sample_notification); err != nil {
				t.Fatalf("Invalid PostApiNotify(): %s", err)
			}
		})
		t.Run("Without subscribers", func(t *testing.T) {
			db.EXPECT().GetTX().Return(sample_tx, nil)
			db.EXPECT().PushNotification(sample_tx, sample_notification).Return(nil)
			db.EXPECT().CommitTX(sample_tx).Return(nil)
			db.EXPECT().GetSubscribers().Return([]models.Subscriber{}, nil)

			if err := bot.Notify(sample_notification); err != nil {
				t.Fatalf("Invalid Notify(): %s", err)
			}
		})
		t.Run("Invalid push", func(t *testing.T) {
			db.EXPECT().GetTX().Return(sample_tx, nil)
			db.EXPECT().PushNotification(sample_tx, sample_notification).Return(fmt.Errorf("some error"))
			db.EXPECT().RollbackTX(sample_tx).Return(nil)

			if err := bot.Notify(sample_notification); err == nil {
				t.Fatalf("Invalid Notify(): error does not exists")
			}
		})
	})

	t.Run("handle()", func(t *testing.T) {
		update := tg.Update{Message: &tg.Message{
			Text: "test",
			Chat: &tg.Chat{ID: 0},
			From: &tg.User{
				UserName:  "test_username",
				FirstName: "test_name",
			},
		}}
		subscriber := models.Subscriber{
			ChatID:   update.Message.Chat.ID,
			Username: update.Message.From.UserName,
			Name:     update.Message.From.FirstName,
		}
		t.Run("New subscriber", func(t *testing.T) {
			db.EXPECT().AddSubscriber(subscriber).Return(false, nil)
			api.EXPECT().Send(tg.NewMessage(update.Message.Chat.ID, "Вы успешно подписаны на оповещения от сервисов сайта Pest Control Expert")).Return(tg.Message{}, nil)
			bot.handle(update)
		})
		t.Run("Exists subscriber", func(t *testing.T) {
			db.EXPECT().AddSubscriber(models.Subscriber{
				ChatID:   update.Message.Chat.ID,
				Username: update.Message.From.UserName,
				Name:     update.Message.From.FirstName,
			}).Return(true, nil)
			api.EXPECT().Send(tg.NewMessage(update.Message.Chat.ID, "Вы уже были подписаны на оповещения от сервисов сайта Pest Control Expert")).Return(tg.Message{}, nil)
			bot.handle(update)
		})
		t.Run("Empty update", func(t *testing.T) {
			bot.handle(tg.Update{})
		})
	})
}
