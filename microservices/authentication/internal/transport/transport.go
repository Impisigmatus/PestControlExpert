package transport

import (
	"github.com/Impisigmatus/PestControlExpert/microservices/authentication/autogen/server"
	"github.com/Impisigmatus/PestControlExpert/microservices/authentication/internal/infrastructure"

	"github.com/go-playground/validator/v10"
)

type handler struct {
	infra    *infrastructure.Infrastructure
	validate *validator.Validate

	issuer    string
	expiresIn int
}

func New(expiresIn int) server.ServerInterface {
	return &handler{
		infra:    infrastructure.New(expiresIn),
		validate: validator.New(),

		issuer:    issuer,
		expiresIn: expiresIn,
	}
}
