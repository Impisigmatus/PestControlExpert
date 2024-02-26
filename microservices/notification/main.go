package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"path"
	"runtime"
	"strconv"
	"strings"
	"syscall"

	"github.com/Impisigmatus/PestControlExpert/microservices/notification/autogen"
	"github.com/Impisigmatus/PestControlExpert/microservices/notification/internal/middlewares"
	"github.com/Impisigmatus/PestControlExpert/microservices/notification/internal/postgres"
	"github.com/Impisigmatus/PestControlExpert/microservices/notification/internal/service"
	"github.com/Impisigmatus/PestControlExpert/microservices/notification/internal/telegram"
	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetReportCaller(true)
	logrus.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
		CallerPrettyfier: func(frame *runtime.Frame) (string, string) {
			file := frame.File[len(path.Dir(os.Args[0]))+1:]
			line := frame.Line
			return "", fmt.Sprintf("%s:%d", file, line)
		},
	})
}

func main() {
	const (
		base       = 10
		size       = 64
		address    = "ADDRESS"
		token      = "TELEGRAM_API_TOKEN"
		password   = "SUBSCRIBE_PASSWORD"
		auth       = "APIS_AUTH_BASIC"
		pgHost     = "POSTGRES_HOSTNAME"
		pgPort     = "POSTGRES_PORT"
		pgDB       = "POSTGRES_DATABASE"
		pgUser     = "POSTGRES_USER"
		pgPassword = "POSTGRES_PASSWORD"
	)

	port, err := strconv.ParseUint(os.Getenv(pgPort), base, size)
	if err != nil {
		logrus.Panicf("Invalid postgres port: %s", err)
	}

	transport := service.NewTransport(
		telegram.NewBot(
			postgres.Config{
				Hostname: os.Getenv(pgHost),
				Port:     port,
				Database: os.Getenv(pgDB),
				User:     os.Getenv(pgUser),
				Password: os.Getenv(pgPassword),
			},
			os.Getenv(token),
			os.Getenv(password),
		),
	)
	router := http.NewServeMux()
	router.Handle("/api/",
		middlewares.Use(middlewares.Use(autogen.Handler(transport),
			middlewares.Authorization(strings.Split(os.Getenv(auth), ","))),
			middlewares.Logger(),
		),
	)

	server := &http.Server{
		Addr:    os.Getenv(address),
		Handler: router,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logrus.Panicf("Invalid service starting: %s", err)
		}
		logrus.Info("Service stopped")
	}()
	logrus.Info("Service started")

	channel := make(chan os.Signal, 1)
	signal.Notify(channel,
		syscall.SIGABRT,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)
	<-channel

	if err := server.Shutdown(context.Background()); err != nil {
		logrus.Panicf("Invalid service stopping: %s", err)
	}
}
