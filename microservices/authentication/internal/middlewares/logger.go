package middlewares

import (
	"net/http"
	"time"

	"github.com/Impisigmatus/service_core/log"
	chi "github.com/go-chi/chi/v5/middleware"
)

type logger struct {
	next http.Handler
}

func Logger() Middleware {
	return func(next http.Handler) http.Handler {
		return &logger{next: next}
	}
}

func (l *logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	handler := chi.RequestLogger(&logger{})
	handler(l.next).ServeHTTP(w, r)
}

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
	log.Infof("%s %s | %s | %s | %d %s", e.Method, e.Path, e.Hostname, duration, status, http.StatusText(status))
}

func (e *entry) Panic(v interface{}, stack []byte) {
	log.Errorf("panic occured panic: %+v stack:%s, ", v, stack)
}
