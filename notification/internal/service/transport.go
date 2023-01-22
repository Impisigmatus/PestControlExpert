package service

import (
	"fmt"
	"io"
	"net/http"

	"github.com/Impisigmatus/PestControlExpert/notification/autogen"
	"github.com/Impisigmatus/PestControlExpert/notification/internal/models"
	"github.com/Impisigmatus/PestControlExpert/notification/internal/telegram"
	"github.com/Impisigmatus/PestControlExpert/notification/internal/utils"
	"github.com/go-playground/validator/v10"
	jsoniter "github.com/json-iterator/go"
)

type Transport struct {
	bot       *telegram.Bot
	validator *validator.Validate
}

func NewTransport(bot *telegram.Bot) autogen.ServerInterface {
	return &Transport{
		bot:       bot,
		validator: validator.New(),
	}
}

func (transport *Transport) PostApiNotify(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		utils.WriteString(w, http.StatusInternalServerError, fmt.Sprintf("Не удалось прочитать тело запроса: %s", err))
		return
	}

	var notification models.Notification
	if err := jsoniter.Unmarshal(data, &notification); err != nil {
		utils.WriteString(w, http.StatusBadRequest, fmt.Sprintf("Не распарсить тело запроса: %s", err))
		return
	}

	if err := transport.validator.Struct(notification); err != nil {
		utils.WriteString(w, http.StatusBadRequest, fmt.Sprintf("Не валидное тело запроса: %s", err))
		return
	}

	if err := transport.bot.Send(notification.Text); err != nil {
		utils.WriteString(w, http.StatusInternalServerError, fmt.Sprintf("Не удалось отправить оповещения: %s", err))
		return
	}

	utils.WriteNoContent(w)
}
