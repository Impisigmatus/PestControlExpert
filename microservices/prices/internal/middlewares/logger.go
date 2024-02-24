package middlewares

import (
	"net/http"
	"time"

	chi "github.com/go-chi/chi/v5/middleware"
	"github.com/sirupsen/logrus"
)

type logger struct{}

func (*logger) NewLogEntry(r *http.Request) chi.LogEntry {
	return &entry{
		Method:   r.Method,
		Hostname: r.RemoteAddr,
		Path:     r.URL.Path,
	}
}

type entry struct {
	Method   string
	Hostname string
	Path     string
}

func (e *entry) Write(status int, _ int, header http.Header, duration time.Duration, extra interface{}) {
	logrus.Infof("%s %s | %s | %s | %d %s", e.Method, e.Path, e.Hostname, duration, status, http.StatusText(status))
}

func (e *entry) Panic(v interface{}, stack []byte) {
	logrus.Errorf("panic occured panic: %+v stack:%s, ", v, stack)
}

type middleware struct {
	next http.Handler
}

func Logger() Middleware {
	return func(next http.Handler) http.Handler {
		return &middleware{next: next}
	}
}

func (ware *middleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	logger := chi.RequestLogger(&logger{})
	logger(ware.next).ServeHTTP(w, r)
}
