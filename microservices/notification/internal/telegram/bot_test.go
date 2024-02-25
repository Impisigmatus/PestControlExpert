package telegram

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/Impisigmatus/PestControlExpert/microservices/notification/autogen"
	"github.com/Impisigmatus/PestControlExpert/microservices/notification/autogen/mocks"
	"github.com/Impisigmatus/PestControlExpert/microservices/notification/internal/models"
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
		sample_subscribers := []models.Subscriber{
			{
				ChatID:   0,
				Username: "user_1",
				Name:     "Василий",
			},
			{
				ChatID:   1,
				Username: "user_2",
				Name:     "Олег",
			},
			{
				ChatID:   2,
				Username: "user_3",
				Name:     "Иван",
			},
		}

		t.Run("SuccessTestCases", func(t *testing.T) {
			t.Run("with subscribers", func(t *testing.T) {
				db.EXPECT().GetTX().Return(sample_tx, nil)
				db.EXPECT().PushNotification(sample_tx, sample_notification).Return(nil)
				db.EXPECT().GetSubscribers().Return(sample_subscribers, nil)
				db.EXPECT().CommitTX(sample_tx).Return(nil)

				api.EXPECT().Send(gomock.Any()).Return(tg.Message{}, nil)
				api.EXPECT().Send(gomock.Any()).Return(tg.Message{}, nil)
				api.EXPECT().Send(gomock.Any()).Return(tg.Message{}, nil)

				if err := bot.Notify(sample_notification); err != nil {
					t.Fatalf("Invalid Notify(): %s", err)
				}
			})
			t.Run("without subscribers", func(t *testing.T) {
				db.EXPECT().GetTX().Return(sample_tx, nil)
				db.EXPECT().PushNotification(sample_tx, sample_notification).Return(nil)
				db.EXPECT().GetSubscribers().Return([]models.Subscriber{}, nil)
				db.EXPECT().CommitTX(sample_tx).Return(nil)

				if err := bot.Notify(sample_notification); err != nil {
					t.Fatalf("Invalid Notify(): %s", err)
				}
			})
		})

		t.Run("FailTestCases", func(t *testing.T) {
			sample_error := fmt.Errorf("some error")
			t.Run("invalid push db", func(t *testing.T) {
				db.EXPECT().GetTX().Return(sample_tx, nil)
				db.EXPECT().PushNotification(sample_tx, sample_notification).Return(sample_error)
				db.EXPECT().RollbackTX(sample_tx).Return(nil)

				if err := bot.Notify(sample_notification); err == nil {
					t.Fatalf("Invalid Notify(): error does not exists")
				}
			})
			t.Run("invalid send tg api", func(t *testing.T) {
				db.EXPECT().GetTX().Return(sample_tx, nil)
				db.EXPECT().PushNotification(sample_tx, sample_notification).Return(nil)
				db.EXPECT().GetSubscribers().Return(nil, sample_error)
				db.EXPECT().RollbackTX(sample_tx).Return(nil)

				if err := bot.Notify(sample_notification); err == nil {
					t.Fatalf("Invalid Notify(): error does not exists")
				}
			})
			t.Run("invalid commit", func(t *testing.T) {
				db.EXPECT().GetTX().Return(sample_tx, nil)
				db.EXPECT().PushNotification(sample_tx, sample_notification).Return(nil)
				db.EXPECT().GetSubscribers().Return(sample_subscribers, nil)
				db.EXPECT().CommitTX(sample_tx).Return(sample_error)
				db.EXPECT().RollbackTX(sample_tx).Return(nil)

				api.EXPECT().Send(gomock.Any()).Return(tg.Message{}, nil)
				api.EXPECT().Send(gomock.Any()).Return(tg.Message{}, nil)
				api.EXPECT().Send(gomock.Any()).Return(tg.Message{}, nil)

				if err := bot.Notify(sample_notification); err == nil {
					t.Fatalf("Invalid Notify(): error does not exists")
				}
			})
		})
	})

	t.Run("handle()", func(t *testing.T) {
		sample_update := tg.Update{Message: &tg.Message{
			Text: "test",
			Chat: &tg.Chat{ID: 0},
			From: &tg.User{
				UserName:  "test_username",
				FirstName: "test_name",
			},
			MessageID: 0,
		}}
		sample_subscriber := models.Subscriber{
			ChatID:   sample_update.Message.Chat.ID,
			Username: sample_update.Message.From.UserName,
			Name:     sample_update.Message.From.FirstName,
		}

		t.Run("SuccessTestCases", func(t *testing.T) {
			t.Run("new subscriber", func(t *testing.T) {
				db.EXPECT().AddSubscriber(sample_subscriber).Return(false, nil)
				api.EXPECT().Send(gomock.Any()).Return(tg.Message{}, nil)
				bot.handle(sample_update)
			})
			t.Run("exists subscriber", func(t *testing.T) {
				db.EXPECT().AddSubscriber(sample_subscriber).Return(true, nil)
				api.EXPECT().Send(gomock.Any()).Return(tg.Message{}, nil)
				bot.handle(sample_update)
			})
			t.Run("empty update", func(t *testing.T) {
				bot.handle(tg.Update{})
			})
		})
		t.Run("FailTestCases", func(t *testing.T) {
			sample_error := fmt.Errorf("some error")
			t.Run("invalid add subscriber", func(t *testing.T) {
				db.EXPECT().AddSubscriber(sample_subscriber).Return(false, sample_error)
				bot.handle(sample_update)
			})
			t.Run("invalid send tg api", func(t *testing.T) {
				db.EXPECT().AddSubscriber(sample_subscriber).Return(false, nil)
				api.EXPECT().Send(gomock.Any()).Return(tg.Message{}, sample_error)
				bot.handle(sample_update)
			})
		})
	})
}
