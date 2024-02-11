package service

import (
	"testing"

	"github.com/Impisigmatus/PestControlExpert/notification/autogen"
	"github.com/Impisigmatus/PestControlExpert/notification/autogen/mock"
	"github.com/Impisigmatus/PestControlExpert/notification/internal/telegram"
	"github.com/golang/mock/gomock"
)

func Test_TelegramBotAPI(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	db := mock.NewMockDatabase(ctrl)
	bot := telegram.NewBot(nil, db, "test")
	t.Run("descr", func(t *testing.T) {
		bot.EXPECT().Notify(autogen.Notification{
			Phone:       "+79991234567",
			Name:        "Ivan",
			Description: nil,
		}).Return(nil)
		if err := bot.Notify(autogen.Notification{}); err != nil {
			t.Fatalf("Invalid PostApiNotify(): %w", err)
		}
	})

	// db := mock.NewMockDatabase(ctrl)
	// db.EXPECT().GetSubscribers().Return([]models.Subscriber{
	// 	{
	// 		ChatID:   0,
	// 		Username: "username",
	// 		Name:     "name",
	// 	},
	// }, nil)
	// db.EXPECT().AddSubscriber(models.Subscriber{
	// 	ChatID:   0,
	// 	Username: "username",
	// 	Name:     "name",
	// }).Return(true, nil)
	// db.EXPECT().AddSubscriber(gomock.Any()).Return(false, nil)
}
