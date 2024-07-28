package infrastructure

import (
	"os"
	"time"

	"github.com/Impisigmatus/PestControlExpert/microservices/authentication/internal/models"
)

type Infrastructure struct {
	secret []byte

	issuer    string
	expiresIn time.Duration
	resources map[string]models.Roles
}

func New(expiresIn int) *Infrastructure {
	return &Infrastructure{
		secret:    []byte(os.Getenv(envtSecret)),
		issuer:    issuer,
		expiresIn: time.Duration(expiresIn) * time.Second,
		resources: map[string]models.Roles{
			resourceHelloWorld: {
				Roles: []string{
					permGet + resourceHelloWorld,
				},
			},
			resourcePingPong: {
				Roles: []string{
					permGet + resourcePingPong,
				},
			},
		},
	}
}
