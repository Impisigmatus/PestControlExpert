package models

type Notification struct {
	Text string `json:"text" validate:"required"`
}
