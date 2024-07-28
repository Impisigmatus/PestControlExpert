package models

// For validating request data
type Payload struct {
	Username string `validate:"required"`
}
