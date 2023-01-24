package database

type Database interface {
	GetSubscribers() ([]int64, error)
	AddSubscriber(id int64) (bool, error)
}
