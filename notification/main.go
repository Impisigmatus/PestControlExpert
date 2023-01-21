package main

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"time"

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

	time.Sleep(10 * time.Second)
	if err := bot.Send("Оповещение"); err != nil {
		logrus.Panicf("Invalid send: %s", err)
	}
}
