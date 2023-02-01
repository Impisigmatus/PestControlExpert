package service

import (
	"fmt"

	"github.com/Impisigmatus/PestControlExpert/notification/internal/models"
	"github.com/Impisigmatus/PestControlExpert/notification/internal/telegram"
)

type Service struct {
	bot *telegram.Bot
}

func NewService(bot *telegram.Bot) *Service {
	return &Service{bot: bot}
}

func (srv *Service) Notify(notification models.Notification) error {
	tx, err := srv.bot.GetTX()
	if err != nil {
		return fmt.Errorf("Invalid transaction: %s", err)
	}

	// TODO: in DB

	if err := srv.bot.Send(notification.Text); err != nil {
		if txErr := tx.Rollback(); txErr != nil {
			err = fmt.Errorf("%s with invalid rollback: %s", err, txErr)
		}

		return fmt.Errorf("Invalid tg bot send: %s", err)
	}

	if err := tx.Commit(); err != nil {
		if txErr := tx.Rollback(); txErr != nil {
			err = fmt.Errorf("%s with invalid rollback: %s", err, txErr)
		}

		return fmt.Errorf("Invalid commit: %s", err)
	}

	return nil
}
