package middlewares

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"

	"github.com/Impisigmatus/PestControlExpert/notification/internal/utils"
)

type authorization struct {
	next    http.Handler
	secrets []string
}

func Authorization(secrets []string) Middleware {
	return func(next http.Handler) http.Handler {
		return &authorization{
			next:    next,
			secrets: secrets,
		}
	}
}

func (auth *authorization) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	const header = "Authorization"

	authorization := r.Header.Get(header)
	if !strings.HasPrefix(authorization, "Basic") {
		utils.WriteString(w, http.StatusUnauthorized, fmt.Errorf("Invalid type"), "Неверный тип авторизации")
		return
	}

	if err := auth.basic(w, authorization); err != nil {
		utils.WriteString(w, http.StatusUnauthorized, err, "Неверные логин или пароль")
		return
	}

	auth.next.ServeHTTP(w, r)
}

func (auth *authorization) basic(w http.ResponseWriter, header string) error {
	const prefix = "Basic "
	authorization := header[len(prefix):]
	data, err := base64.StdEncoding.DecodeString(authorization)
	if err != nil {
		return fmt.Errorf("Invalid decode basic authorization: %s", err)
	}
	decoded := string(data)

	for _, secret := range auth.secrets {
		if decoded == secret {
			return nil
		}
	}

	return fmt.Errorf("Invalid basic authorization")
}
