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
	if err := h.authorize(r); err != nil {
		utils.WriteString(w, http.StatusUnauthorized, err, "authorization failed")
		return
	}

	utils.WriteNoContent(w)
}

func (h *handler) authorize(r *http.Request) error {
	auth := r.Header.Get(headerAuthorization)
	if len(auth) < skipLength {
		return AuthorizationHeaderError
	}

	claims, err := h.infra.GetTokenClaims(auth[skipLength-1:])
	if err != nil {
		return err
	}

	_ = claims

	return nil
}
