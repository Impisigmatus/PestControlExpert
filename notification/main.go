package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"path"
	"runtime"
	"strings"
	"syscall"

	"github.com/Impisigmatus/PestControlExpert/notification/autogen"
	"github.com/Impisigmatus/PestControlExpert/notification/internal/middlewares"
	"github.com/Impisigmatus/PestControlExpert/notification/internal/service"
	"github.com/Impisigmatus/PestControlExpert/notification/internal/telegram"
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
		token    = "PCE_TELEGRAM_API_TOKEN"
		password = "PCE_SUBSCRIBE_PASSWORD"
		auth     = "APIS_AUTH_BASIC"
	)
	bot := telegram.NewBot(os.Getenv(token), os.Getenv(password))

	transport := service.NewTransport(bot)
	router := http.NewServeMux()
	router.Handle("/api/",
		middlewares.Use(middlewares.Use(autogen.Handler(transport),
			middlewares.Authorization(strings.Split(os.Getenv(auth), ","))),
			middlewares.Logger(),
		),
	)

	const addr = ":8000"
	server := &http.Server{
		Addr:    addr,
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
