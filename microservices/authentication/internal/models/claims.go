package models

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	jwt.RegisteredClaims

	Username  string    `json:"username"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiresAt time.Time `json:"expires_at"`
	ExpiresIn int       `json:"expires_in"`
	Type      string    `json:"type"`

	Resources map[string]Roles `json:"resource_access"`
}

type Roles struct {
	Roles []string `json:"roles"`
}
