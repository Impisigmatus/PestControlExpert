package middlewares

import "errors"

var (
	AuthorizationHeaderError = errors.New("invalid Authorization header")
	AuthorizationMethodError = errors.New("invalid authorization method")
	BasicAuthorizationError  = errors.New("invalid basic authorization")
)
