package middlewares

import "errors"

var AuthorizationHeaderError = errors.New("invalid Authorization header")
