package transport

import (
	"net/http"

	"github.com/Impisigmatus/PestControlExpert/microservices/authentication/internal/models"
	"github.com/Impisigmatus/service_core/utils"
)

// Set godoc
//
// @Router /login [get]
// @Summary Получение токена
// @Description Генерирует JWT токен
//
// @Tags Authentication
//
// @Produce application/json
//
// @Success 200 {object} token_response "Запрос выполнен успешно"
// @Failure 400 {object} nil "Ошибка валидации данных"
// @Failure 401 {object} nil "Ошибка авторизации"
// @Failure 500 {object} nil "Произошла внутренняя ошибка сервера"
func (h *handler) GetLogin(w http.ResponseWriter, r *http.Request) {
	username, _, _ := r.BasicAuth()
	payload := models.Payload{
		Username: username,
	}
	if err := h.validate.Struct(payload); err != nil {
		utils.WriteString(w, http.StatusBadRequest, err, "invalid request")
		return
	}

	claims := h.infra.GetClaims(payload.Username)
	signed, err := h.infra.GetSignedToken(claims)
	if err != nil {
		utils.WriteString(w, http.StatusInternalServerError, err, "invalid token")
		return
	}

	utils.WriteObject(w, &models.TokenResponse{
		AccessToken: signed,
		ExpiresIn:   h.expiresIn,
		TokenType:   tokenType,
	})
}
