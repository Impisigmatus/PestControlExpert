package main

import (
	"fmt"
	"net/http"
	"os"
	"path"
	"runtime"

	"github.com/Impisigmatus/PestControlExpert/notification/autogen"
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
	)
	bot := telegram.NewBot(os.Getenv(token), os.Getenv(password))

	transport := service.NewTransport(bot)
	router := http.NewServeMux()
	router.Handle("/api/", autogen.Handler(transport))

	const addr = ":8000"
	server := &http.Server{
		Addr:    addr,
		Handler: router,
	}

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logrus.Panic(err)
	}
}
