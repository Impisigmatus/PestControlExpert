package main

import (
	"net/http"

	_ "github.com/Impisigmatus/PestControlExpert/microservices/authentication/autogen/docs"
	"github.com/Impisigmatus/PestControlExpert/microservices/authentication/autogen/server"
	"github.com/Impisigmatus/PestControlExpert/microservices/authentication/internal/infrastructure"
	"github.com/Impisigmatus/PestControlExpert/microservices/authentication/internal/middlewares"
	"github.com/Impisigmatus/PestControlExpert/microservices/authentication/internal/transport"
	"github.com/Impisigmatus/service_core/log"
	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Authentication API
// @version 1.0
// @description %README_FILE%
// @host :8000
// @BasePath /
func main() {
	log.Init(log.LevelDebug)

	infra := infrastructure.New(18000)

	router := chi.NewRouter()
	router.Handle("/*",
		middlewares.Use(
			middlewares.Use(
				server.Handler(transport.New(infra)),
				middlewares.Logger(),
			),
			middlewares.Authorization(infra),
		),
	)
	router.Get("/swagger/*", httpSwagger.Handler())

	server := &http.Server{
		Addr:    ":8000",
		Handler: router,
	}

	log.Info("service started")
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Panic("failed to start service", err)
	}
}
