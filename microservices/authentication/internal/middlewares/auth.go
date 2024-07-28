package middlewares

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"

	"github.com/Impisigmatus/service_core/log"
	"github.com/Impisigmatus/service_core/utils"

	"github.com/Impisigmatus/PestControlExpert/microservices/authentication/internal/infrastructure"
)

type authorization struct {
	next  http.Handler
	infra *infrastructure.Infrastructure
}

func Authorization(infra *infrastructure.Infrastructure) Middleware {
	return func(next http.Handler) http.Handler {
		return &authorization{
			next:  next,
			infra: infra,
		}
	}
}

func (auth *authorization) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	data := r.Header.Get(headerAuthorization)
	if len(data) < skipLength {
		utils.WriteString(w, http.StatusUnauthorized, nil, "invalid authorization header")
		return
	}

	switch {
	case strings.HasPrefix(data, "Basic"):
		if err := auth.basic("Basic:", data); err != nil {
			utils.WriteString(w, http.StatusUnauthorized, err, "authorization failed")
			return
		}
	case strings.HasPrefix(data, "Bearer"):
		if err := auth.token(data); err != nil {
			utils.WriteString(w, http.StatusUnauthorized, err, "authorization failed")
			return
		}
	default:
		utils.WriteString(w, http.StatusUnauthorized, fmt.Errorf("invalid authorization method"), "authorization failed")
		return
	}

	auth.next.ServeHTTP(w, r)
}

func (auth *authorization) token(data string) error {
	claims, err := auth.infra.GetTokenClaims(data[skipLength-1:])
	if err != nil {
		return err
	}

	log.Infof("claims: %v", *claims)
	return nil
}

func (auth *authorization) basic(prefix string, data string) error {
	authorization := data[len(prefix):]
	decoded, err := base64.StdEncoding.DecodeString(authorization)
	if err != nil {
		return fmt.Errorf("Invalid decode basic authorization: %s", err)
	}

	if string(decoded) != "dev:test" {
		return fmt.Errorf("Invalid basic authorization")
	}

	return nil
}
