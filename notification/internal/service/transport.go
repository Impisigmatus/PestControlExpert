package service

import (
	"fmt"
	"io"
	"net/http"

	"github.com/Impisigmatus/PestControlExpert/notification/autogen"
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
		utils.WriteString(w, http.StatusInternalServerError, fmt.Errorf("Invalid read body: %s", err), "Неудалось прочитать тело запроса")
		return
	}

	var notification autogen.Notification
	if err := jsoniter.Unmarshal(data, &notification); err != nil {
		utils.WriteString(w, http.StatusBadRequest, fmt.Errorf("Invalid parse body: %s", err), "Невалидное удалось распарсить тело запроса формата JSON")
		return
	}

	if err := transport.validator.Struct(notification); err != nil {
		utils.WriteString(w, http.StatusBadRequest, fmt.Errorf("Invalid body: %s", err), "Невалидное тело запроса")
		return
	}

	if notification.Description != nil && len(*notification.Description) == 0 {
		notification.Description = nil
	}

	if err := transport.bot.Notify(notification); err != nil {
		utils.WriteString(w, http.StatusInternalServerError, fmt.Errorf("Invalid notify: %s", err), "Неудалось отправить оповещения")
		return
	}

	utils.WriteNoContent(w)
}
