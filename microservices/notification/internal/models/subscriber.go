package models

type Subscriber struct {
	ChatID   int64  `db:"chat_id" validate:"required"`
	Username string `db:"username" validate:"required"`
	Name     string `db:"name" validate:"required"`
}
