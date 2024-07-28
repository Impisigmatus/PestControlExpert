package middlewares

import "net/http"

type Middleware = func(http.Handler) http.Handler

func Use(handler http.Handler, middleware Middleware) http.Handler {
	return middleware(handler)
}
