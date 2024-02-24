package utils

import (
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"
)

func WriteNoContent(w http.ResponseWriter) {
	WriteString(w, http.StatusNoContent, nil, "")
}

func WriteString(w http.ResponseWriter, status int, err error, str string, args ...any) {
	const (
		header  = "Content-Type"
		content = "text/plain"
	)

	if err != nil {
		logrus.Error(err)
	}

	w.Header().Set(header, content)
	w.WriteHeader(status)

	if _, err := w.Write([]byte(fmt.Sprintf(str, args...))); err != nil {
		logrus.Errorf("Invalid write response body: %s", err)
	}
}
