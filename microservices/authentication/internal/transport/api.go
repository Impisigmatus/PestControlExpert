package transport

import (
	"net/http"

	"github.com/Impisigmatus/service_core/utils"
)

// Set godoc
//
// @Router /api [get]
// @Summary Тестовое API
// @Description Данное API предназначено для проверки прав пользователя
//
// @Tags APIs
//
// @Success 204 {object} nil "Запрос выполнен успешно"
// @Failure 400 {object} nil "Ошибка валидации данных"
// @Failure 401 {object} nil "Ошибка авторизации"
// @Failure 500 {object} nil "Произошла внутренняя ошибка сервера"
func (h *handler) GetApi(w http.ResponseWriter, r *http.Request) {
	utils.WriteNoContent(w)
}
