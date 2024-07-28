package middlewares

import (
	"encoding/base64"
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
		utils.WriteString(w, http.StatusUnauthorized, AuthorizationHeaderError, "invalid header")
		return
	}

	switch {
	case strings.HasPrefix(data, prefixBasic):
		if err := auth.basic(data); err != nil {
			utils.WriteString(w, http.StatusUnauthorized, err, "authorization failed")
			return
		}
	case strings.HasPrefix(data, "Bearer"):
		if err := auth.token(data); err != nil {
			utils.WriteString(w, http.StatusUnauthorized, err, "authorization failed")
			return
		}
	default:
		utils.WriteString(w, http.StatusUnauthorized, AuthorizationMethodError, "authorization failed")
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

func (auth *authorization) basic(data string) error {
	authorization := data[len(prefixBasic)+1:]
	decoded, err := base64.StdEncoding.DecodeString(authorization)
	if err != nil {
		return err
	}

	if string(decoded) != secret {
		return BasicAuthorizationError
	}

	return nil
}
