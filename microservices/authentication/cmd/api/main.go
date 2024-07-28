package main

import (
	"net/http"

	_ "github.com/Impisigmatus/PestControlExpert/microservices/authentication/autogen/docs"
	"github.com/Impisigmatus/PestControlExpert/microservices/authentication/autogen/server"
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

	router := chi.NewRouter()
	router.Handle("/*", server.Handler(transport.New(18000)))
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
